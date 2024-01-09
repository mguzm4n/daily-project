package main

import (
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
