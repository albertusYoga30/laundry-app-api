package repository

import (
	"database/sql"
)

func CheckUserById(db *sql.DB, id int) (exist bool, err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1 AND role_id=1", id).Scan(&count)
	if err != nil {
		return false, err
	}
	exist = count > 0
	return exist, nil
}
