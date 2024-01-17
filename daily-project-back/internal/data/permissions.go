package data

import "database/sql"

// user.Roles

type PermissionModel struct {
	DB *sql.DB
}

type Permission struct {
	PermissionID int64
	Code         string
}

func (m PermissionModel) GetAllForRole(roleId int64) ([]Permission, error) {
	query := `
		SELECT permissions.id, permissions.code
		FROM roles_permissions, permissions
		WHERE roles_permissions.permission_id = permissions.id 
			AND roles.id = $1
	`

	rows, err := m.DB.Query(query, roleId)
	if err != nil {
		return nil, err
	}

	var permissions []Permission
	for rows.Next() {
		var permission Permission
		err := rows.Scan(&permission.PermissionID, &permission.Code)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (m PermissionModel) GetAllForUser(userId int64) ([]Permission, error) {
	query := `
	SELECT DISTINCT permissions.id, permissions.code
	FROM roles_users, roles_permissions, permissions
	WHERE roles_users.user_id = $1
		AND roles_users.role_id = roles_permissions.role_id
		AND roles_permissions.permission_id = permissions.id
	
	`

	rows, err := m.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var permissions []Permission
	for rows.Next() {
		var permission Permission
		err := rows.Scan(&permission.PermissionID, &permission.Code)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
