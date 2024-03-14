package model

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	RoleID      int    `json:"role_id"`
}

type Role struct {
	ID   int    `json:"role_id"`
	Name string `json:"role_name"`
	Desc string `json:"role_desc"`
}
