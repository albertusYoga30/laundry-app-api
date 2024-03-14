package repository

import (
	"database/sql"
	"laundry-app-api/model"
)

func GetAllServices(db *sql.DB) (result []model.Service, err error) {

	rows, err := db.Query("SELECT * FROM laundry_services")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data = model.Service{}
		err = rows.Scan(&data.ServiceID, &data.ServiceName, &data.ServiceDesc, &data.ServicePrice, &data.UpdatedAt, &data.CreatedAt)
		if err != nil {
			panic(err)
		}
		result = append(result, data)
	}
	return
}

func CheckServiceData(db *sql.DB, id int) (exist bool, err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM laundry_services WHERE service_id = $1", id).Scan(&count)
	if err != nil {
		return false, err
	}
	exist = count > 0
	return exist, nil
}

func CheckServiceName(db *sql.DB, name string) (exist bool, err error) {
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM laundry_services WHERE service_name = $1", name).Scan(&count)
	if err != nil {
		return false, err
	}
	exist = count > 0
	return exist, nil
}

func GetServicePrice(db *sql.DB, id int) (price int, err error) {
	var data model.Service
	err = db.QueryRow("SELECT service_price FROM laundry_services WHERE service_id=$1", id).Scan(&data.ServicePrice)
	if err != nil {
		return 0, err
	}
	return data.ServicePrice, nil
}

func InsertService(db *sql.DB, service model.Service) (err error) {
	sql := "INSERT INTO laundry_services (service_name, service_desc, service_price) VALUES($1,$2,$3)"
	errs := db.QueryRow(sql, service.ServiceName, service.ServiceDesc, service.ServicePrice)
	return errs.Err()
}

func UpdateService(db *sql.DB, service model.Service) (err error) {
	sql := "UPDATE laundry_services SET service_name=$1, service_desc=$2, service_price=$3 WHERE service_id=$4"
	errs := db.QueryRow(sql, service.ServiceName, service.ServiceDesc, service.ServicePrice, service.ServiceID)
	return errs.Err()
}

func DeleteService(db *sql.DB, service model.Service) (err error) {
	errs := db.QueryRow("DELETE FROM laundry_services WHERE service_id=$1", service.ServiceID)
	return errs.Err()
}
