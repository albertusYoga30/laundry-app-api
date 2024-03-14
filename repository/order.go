package repository

import (
	"database/sql"
	"laundry-app-api/model"
)

func GetAllOrders(db *sql.DB) (result []model.Order, err error) {

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data = model.Order{}
		err = rows.Scan(&data.OrderID, &data.UserID, &data.Quantity, &data.TotalPrice, &data.ServiceID, &data.DurationID, &data.CreatedAt)
		if err != nil {
			panic(err)
		}
		result = append(result, data)
	}
	return
}

func GetOrdersByUserId(db *sql.DB, userId int) (result []model.Order, err error) {
	rows, err := db.Query("SELECT * FROM orders WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data = model.Order{}
		err = rows.Scan(&data.OrderID, &data.UserID, &data.Quantity, &data.TotalPrice, &data.ServiceID, &data.DurationID, &data.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return
}

func InsertOrder(db *sql.DB, order model.Order) (err error) {
	sql := "INSERT INTO orders (user_id, quantity, total_price, service_id, duration_id) VALUES($1, $2, $3, $4, $5)"
	errs := db.QueryRow(sql, order.UserID, order.Quantity, order.TotalPrice, order.ServiceID, order.DurationID)
	return errs.Err()
}
