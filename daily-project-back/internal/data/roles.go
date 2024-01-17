package data

import "database/sql"

// user.Roles

type RoleModel struct {
	DB *sql.DB
}

type Role struct {
	RoleID int64
	Name   string
}

func (m RoleModel) Get() ([]Role, error) {
	query := `
		SELECT roles.id, roles.name
		FROM roles
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var roles []Role
	for rows.Next() {
		var role Role
		err := rows.Scan(&role.RoleID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (m RoleModel) GetAllForUser(userId int64) ([]Role, error) {
	query := `
		SELECT roles.id, roles.name
		FROM roles_user, roles
		WHERE roles_user.role_id = roles.id 
			AND roles_user.user_id = $1
	`

	rows, err := m.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var roles []Role
	for rows.Next() {
		var role Role
		err := rows.Scan(&role.RoleID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}
