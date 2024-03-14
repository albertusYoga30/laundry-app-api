package main

import (
	"database/sql"
	"fmt"
	"laundry-app-api/controllers"
	"laundry-app-api/database"
	"laundry-app-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	router.POST("/register/admin", controllers.RegisterAdmin)
	router.POST("/register/customer", controllers.RegisterCustomer)
	router.POST("/login", controllers.Login)

	router.GET("/orders", middleware.JWTAuth(2), controllers.GetAllOrders)
	router.GET("/orders/:id", middleware.JWTAuth(1), controllers.GetAllOrdersByUserId)
	router.POST("/orders", middleware.JWTAuth(1), controllers.InsertOrder)

	router.GET("/services", controllers.GetAllServices)
	router.POST("/service", middleware.JWTAuth(2), controllers.InsertService)
	router.PUT("/service/:id", middleware.JWTAuth(2), controllers.UpdateService)
	router.DELETE("/service/:id", middleware.JWTAuth(2), controllers.DeleteService)

	router.GET("/durations", controllers.GetAllDurations)
	router.POST("/duration", middleware.JWTAuth(2), controllers.InsertDuration)
	router.PUT("/duration/:id", middleware.JWTAuth(2), controllers.UpdateDuration)
	router.DELETE("/duration/:id", middleware.JWTAuth(2), controllers.DeleteDuration)

	router.Run("localhost:8080")

}
