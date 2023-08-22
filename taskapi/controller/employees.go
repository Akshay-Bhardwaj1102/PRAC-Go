package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskapi.com/config"
	"taskapi.com/model"
)

// function to find all leaves

func FindAll(c *gin.Context) {

	var emp []model.Employee
	config.DB.Find(&emp)
	c.JSON(http.StatusOK, gin.H{"data": emp})
}

// function to create a leave

func CreateLeave(c *gin.Context) {

	var input model.Employee

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new task
	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})

}

// function to find latest leave by name

func FindLeave(c *gin.Context) {
	var ename model.Employee
	if err := config.DB.Where("name = ?", c.Param("name")).First(&ename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ename})
}

// function to find all leave by ename

func FindAllByName(c *gin.Context) {
	var ename []model.Employee
	if err := config.DB.Where("name = ?", c.Param("name")).Find(&ename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found !"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ename})
}

//function to delete leave {by name}

func DeleteLeave(c *gin.Context) {
	var ename model.Employee

	if err := config.DB.Where("name = ?", c.Param("name")).First(&ename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&ename)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
