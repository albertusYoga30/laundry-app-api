package controllers

import (
	"laundry-app-api/database"
	"laundry-app-api/model"
	"laundry-app-api/repository"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func RegisterCustomer(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, phone_number, err := repository.ValidateUser(database.DbConnection, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist!"})
		return
	}

	if email || phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Phone number already Used!"})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	user.RoleID = 1

	err = repository.Register(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "User registered successfully!",
	})
}

func RegisterAdmin(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, email, phone_number, err := repository.ValidateUser(database.DbConnection, user.Username, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exist!"})
		return
	}

	if email || phone_number {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Phone number already Used!"})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPass)
	user.RoleID = 2

	err = repository.Register(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "User registered successfully!",
	})
}

func Login(c *gin.Context) {

	var credential model.Login

	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUserByUsername(database.DbConnection, credential.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))
	if errPass != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"role":      user.RoleID,
		"ExpiresAt": time.Now().Add(5 * time.Minute).Unix(),
	})

	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}
