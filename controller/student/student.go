package controller

import (
	"net/http"
	"strconv"

	"github.com/Hasyim-Kai/Go_Gin_Rest_Api_MySkill/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostHandler(c *gin.Context, db *gorm.DB) {
	var Student model.Student
	if c.Bind(&Student) == nil {
		result := db.Create(&Student) // pass pointer of data to Create
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success create"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
}

func GetAllHandler(c *gin.Context, db *gorm.DB) {
	var Student []model.Student
	result := db.Find(&Student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}
	if Student == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Student})
}

func GetHandler(c *gin.Context, db *gorm.DB) {
	var Student model.Student
	id, parseErr := strconv.Atoi(c.Param("id"))
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Param ID Must be a Number"})
		return
	}

	Student.Student_id = id
	result := db.First(&Student)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Student})
}

func PutHandler(c *gin.Context, db *gorm.DB) {
	var Student model.Student
	id, parseErr := strconv.Atoi(c.Param("id"))
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Param ID Must be a Number"})
		return
	}

	Student.Student_id = id
	if c.Bind(&Student) == nil {
		result := db.Save(&Student)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Success Update"})
	}

}

func DelHandler(c *gin.Context, db *gorm.DB) {
	var Student model.Student
	id, parseErr := strconv.Atoi(c.Param("id"))
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Param ID Must be a Number"})
		return
	}

	Student.Student_id = id
	result := db.Delete(&Student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success Delete"})
}
