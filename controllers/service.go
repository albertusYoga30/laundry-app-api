package controllers

import (
	"laundry-app-api/database"
	"laundry-app-api/model"
	"laundry-app-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllServices(c *gin.Context) {
	var result gin.H

	laundries, err := repository.GetAllServices(database.DbConnection)
	if err != nil {
		result = gin.H{
			"error": err,
		}
	} else {
		result = gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"result":  laundries,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertService(c *gin.Context) {
	var service model.Service

	err := c.ShouldBindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err = repository.InsertService(database.DbConnection, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Insert Service into Database!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Service!"})
}

func UpdateService(c *gin.Context) {
	var service model.Service
	id, errStr := strconv.Atoi(c.Param("id"))
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	service.ServiceID = id

	dataExist, errData := repository.CheckServiceData(database.DbConnection, service.ServiceID)
	if errData != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errData.Error()})
		return
	}

	if !dataExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data isn't exist!"})
		return
	}

	err := c.ShouldBindJSON(&service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err = repository.UpdateService(database.DbConnection, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Update Service into Database!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Service!"})
}

func DeleteService(c *gin.Context) {
	var service model.Service

	id, errStr := strconv.Atoi(c.Param("id"))
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	service.ServiceID = id

	dataExist, errData := repository.CheckServiceData(database.DbConnection, service.ServiceID)
	if errData != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errData.Error()})
		return
	}

	if !dataExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data isn't exist!"})
		return
	}

	err := repository.DeleteService(database.DbConnection, service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Update Service into Database!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Service"})
}
