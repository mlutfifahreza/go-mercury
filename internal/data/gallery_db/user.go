package gallery_db

import (
	"database/sql"
	"errors"
	"strings"

	"go-mercury/pkg/constant"
)

func (d *DB) CreateUserTab(user User) error {
	db, err := d.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO users (username, password_hash, roles)
		VALUES ($1, $2, $3)`
	err = db.QueryRow(sqlStatement, user.Username, user.PasswordHash, convertRoleToDB(user.Roles)).Err()
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetUserTab(username string) (*User, error) {
	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT username, password_hash, roles FROM users WHERE username = $1`
	var user User
	var roles string

	err = db.QueryRow(sqlStatement, username).Scan(&user.Username, &user.PasswordHash, &roles)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &user, constant.UserNotFoundError
		}
		return &user, err
	}

	user.Roles = convertRoleFromDB(roles)

	return &user, nil
}

func convertRoleFromDB(roles string) []UserRole {
	roleList := strings.Split(roles, ",")
	var result []UserRole
	for _, role := range roleList {
		result = append(result, UserRole(role))
	}

	return result
}

func convertRoleToDB(roles []UserRole) string {
	var rolesStringList []string
	for _, r := range roles {
		rolesStringList = append(rolesStringList, string(r))
	}

	return strings.Join(rolesStringList, ",")
}
