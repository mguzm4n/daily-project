package main

import (
	"errors"
	"strings"

	"daily-project/internal/data"
	"daily-project/internal/validator"

	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.Header().Set("Connection", "close")
					app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
				}
			}()

			next.ServeHTTP(w, r)
		},
	)
}

func (app *application) rateLimit(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mux     sync.Mutex
		clients = make(map[string]*client)
	)

	go func() {
		for {
			time.Sleep(1 * time.Minute)
			mux.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) >= 3*time.Minute {
					delete(clients, ip)
				}
			}
			mux.Unlock()
		}
	}()

	// Initialize new rate limiter.
	reqsPerSecond := 2.0 // r
	maxReqsOnBurst := 4  // b

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if app.config.limiter.enabled {
				ip, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					app.serverErrorResponse(w, r, err)
					return
				}

				// Prevent race conditions from here
				mux.Lock()
				if _, found := clients[ip]; !found {
					clients[ip] = &client{
						limiter: rate.NewLimiter(rate.Limit(reqsPerSecond), maxReqsOnBurst),
					}
				}
				clients[ip].lastSeen = time.Now()

				// Call Allow() on the current limiter for specific ip
				if !clients[ip].limiter.Allow() {
					mux.Unlock()
					app.rateLimitExceededResponse(w, r)
					return
				}

				// Unlock the mutex so clients map is writtable again.
				// Do not use "defer" since that would mean mutex wouldn't be
				// unlocked till all handlers returned.
				mux.Unlock()
			}

			next.ServeHTTP(w, r)
		},
	)
}

func (app *application) authenticate(next http.Handler) http.Handler {
	middlewareFunc := func(w http.ResponseWriter, r *http.Request) {
		// Response may vary depending on value of this Auth header.
		w.Header().Add("Vary", "Authorization")

		// get the authorization header
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			r = app.contextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		// expect format "Bearer + " " + <Token>"
		headerSplit := strings.Split(authorizationHeader, " ")
		if len(headerSplit) != 2 || headerSplit[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerSplit[1]
		v := validator.New()
		if data.ValidateTokenPlaintext(v, token); !v.Valid() {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// get details of user knowing it's valid
		user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidAuthenticationTokenResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		r = app.contextSetUser(r, user)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(middlewareFunc)
}

func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	middlewareFunc := func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)
		if user.IsAnonymous() {
			app.authenticationRequiredResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(middlewareFunc)
}

func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	middlewareFunc := func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		if !user.Activated {
			app.inactiveAccountResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}

	return app.requireAuthenticatedUser(middlewareFunc)
}

func (app *application) requirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)
		permissions, err := app.models.Permissions.GetAllForUser(user.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if !permissions.Includes(code) {
			app.notPermittedResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	}

	return app.requireActivatedUser(fn)
}
