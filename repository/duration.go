package repository

import (
	"database/sql"
	"laundry-app-api/model"
)

func CheckDurationData(db *sql.DB, id int) (exist bool, err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM durations WHERE duration_id = $1", id).Scan(&count)
	if err != nil {
		return false, err
	}
	exist = count > 0
	return exist, nil
}

func GetAllDurations(db *sql.DB) (result []model.Duration, err error) {
	rows, err := db.Query("SELECT * FROM durations")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data = model.Duration{}
		err = rows.Scan(&data.DurationID, &data.DurationName, &data.DurationDays, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			panic(err)
		}
		result = append(result, data)
	}
	return
}

func GetDurationDays(db *sql.DB, id int) (days int, err error) {
	err = db.QueryRow("SELECT duration_days FROM durations WHERE duration_id = $1", id).Scan(&days)
	if err != nil {
		return 0, err
	}
	return days, nil
}

func InsertDuration(db *sql.DB, duration model.Duration) (err error) {
	sql := "INSERT INTO durations (duration_name, duration_days) VALUES($1,$2)"
	errs := db.QueryRow(sql, duration.DurationName, duration.DurationDays)
	return errs.Err()
}

func UpdateDuration(db *sql.DB, duration model.Duration) (err error) {
	sql := "UPDATE durations SET duration_name=$1, duration_days=$2 WHERE duration_id=$3"
	errs := db.QueryRow(sql, duration.DurationName, duration.DurationDays, duration.DurationID)
	return errs.Err()
}

func DeleteDuration(db *sql.DB, duration model.Duration) (err error) {
	sql := "DELETE FROM durations WHERE duration_id=$1"
	errs := db.QueryRow(sql, duration.DurationID)
	return errs.Err()
}
