package repository

import (
	"database/sql"
	"errors"
	"laundry-app-api/model"
)

func Register(db *sql.DB, user model.User) (err error) {
	sql := "INSERT INTO users(username, password, email, address, phone_number,role_id) VALUES($1, $2, $3, $4, $5, $6)"
	errs := db.QueryRow(sql, user.Username, user.Password, user.Email, user.Address, user.PhoneNumber, user.RoleID)
	return errs.Err()
}

func ValidateUser(db *sql.DB, username, email, phone string) (usernameExists, emailExits, phoneExists bool, err error) {

	var usernameCount, emailCount, phoneCount int

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE Username = $1", username).Scan(&usernameCount)
	if err != nil {
		return false, false, false, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&emailCount)
	if err != nil {
		return false, false, false, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = $1", phone).Scan(&phoneCount)
	if err != nil {
		return false, false, false, err
	}

	usernameExists = usernameCount > 0
	emailExits = emailCount > 0
	phoneExists = phoneCount > 0

	return usernameExists, emailExits, phoneExists, nil
}

func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {
	query := "SELECT id, username, password ,role_id FROM users WHERE username = $1"
	user := &model.User{}
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.RoleID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
