package controllers

import (
	"laundry-app-api/database"
	"laundry-app-api/model"
	"laundry-app-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var result gin.H

	orders, err := repository.GetAllOrders(database.DbConnection)
	if err != nil {
		result = gin.H{
			"error": err,
		}
	} else {
		result = gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"result":  orders,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAllOrdersByUserId(c *gin.Context) {
	var result gin.H
	id, errStr := strconv.Atoi(c.Param("id"))
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	orders, err := repository.GetOrdersByUserId(database.DbConnection, id)
	if err != nil {
		result = gin.H{
			"error": err,
		}
	} else {
		result = gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"result":  orders,
		}
	}

	c.JSON(http.StatusOK, result)

}

func InsertOrder(c *gin.Context) {
	var order model.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	user, errUser := repository.CheckUserById(database.DbConnection, order.UserID)
	if errUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errUser.Error()})
		return
	}
	if !user {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User isn't exist!"})
		return
	}

	service, errService := repository.CheckServiceData(database.DbConnection, order.ServiceID)
	if errService != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errService.Error()})
		return
	}
	if !service {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Service isn't exist!"})
		return
	}
	price, errPrice := repository.GetServicePrice(database.DbConnection, order.ServiceID)
	if errPrice != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errPrice.Error()})
		return
	}

	duration, errDuration := repository.CheckDurationData(database.DbConnection, order.DurationID)
	if errDuration != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDuration.Error()})
		return
	}
	if !duration {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Duration isn't exist!"})
		return
	}

	days, errDays := repository.GetDurationDays(database.DbConnection, order.DurationID)
	if errDays != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errDays.Error()})
		return
	}

	if order.Quantity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity Shouldn't be 0!"})
		return
	}

	switch days {
	case 1:
		order.TotalPrice = (price * order.Quantity) * 2
	case 2:
		order.TotalPrice = (price * order.Quantity) + ((price * order.Quantity) / 2)
	default:
		order.TotalPrice = (price * order.Quantity)
	}

	err = repository.InsertOrder(database.DbConnection, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Order!"})
}
