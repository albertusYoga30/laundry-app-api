package controllers

import (
	"laundry-app-api/database"
	"laundry-app-api/model"
	"laundry-app-api/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDurations(c *gin.Context) {
	var result gin.H

	durations, err := repository.GetAllDurations(database.DbConnection)
	if err != nil {
		result = gin.H{"error": err}
	} else {
		result = gin.H{
			"code":    http.StatusOK,
			"message": "Success",
			"result":  durations,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertDuration(c *gin.Context) {
	var duration model.Duration

	err := c.ShouldBindJSON(&duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	dataExist, errData := repository.CheckDurationData(database.DbConnection, duration.DurationID)
	if errData != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errData.Error()})
		return
	}

	if dataExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data already exist!"})
		return
	}

	nameExist, errName := repository.CheckDurationName(database.DbConnection, duration.DurationName)
	if errName != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errName.Error()})
		return
	}

	if nameExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data already exist!"})
		return
	}

	err = repository.InsertDuration(database.DbConnection, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Insert Duration into Database!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success Insert Duration!"})
}

func UpdateDuration(c *gin.Context) {
	var duration model.Duration
	id, errStr := strconv.Atoi(c.Param("id"))
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	duration.DurationID = id

	dataExist, errData := repository.CheckDurationData(database.DbConnection, duration.DurationID)
	if errData != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errData.Error()})
		return
	}

	if !dataExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data isn't exist!"})
		return
	}

	err := c.ShouldBindJSON(&duration)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err = repository.UpdateDuration(database.DbConnection, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Duration"})
}

func DeleteDuration(c *gin.Context) {
	var duration model.Duration

	id, errStr := strconv.Atoi(c.Param("id"))
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	duration.DurationID = id

	dataExist, errData := repository.CheckDurationData(database.DbConnection, duration.DurationID)
	if errData != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errData.Error()})
		return
	}

	if !dataExist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data isn't exist!"})
		return
	}

	err := repository.DeleteDuration(database.DbConnection, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Update Duration into Database!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success Update Duration"})
}
