package controller

import (
	"net/http"

	"github.com/Hasyim-Kai/Go_Gin_Rest_Api/model"
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
	}

	if Student == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Student})

}

// func GetHandler(c *gin.Context, db *gorm.DB) {
// 	var Student []model.Student

// 	studentId := c.Param("student_id")

// 	row, err := db.Query("select * from students where student_id = $1", studentId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	}

// 	utils.RowToStruct(row, &Student)

// 	if Student == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": Student})
// }

// func PutHandler(c *gin.Context, db *gorm.DB) {
// 	var Student model.Student

// 	studentId := c.Param("student_id")

// 	if c.Bind(&Student) == nil {
// 		_, err := db.Exec("update students set student_name=$1 where student_id=$2", Student.Student_name, studentId)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "success update"})
// 	}

// }

// func DelHandler(c *gin.Context, db *gorm.DB) {
// 	studentId := c.Param("student_id")

// 	_, err := db.Exec("delete from students where student_id=$1", studentId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "success delete"})

// }
