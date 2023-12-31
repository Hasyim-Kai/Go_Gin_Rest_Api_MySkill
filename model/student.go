package model

type Student struct {
	Student_id       int    `json:"student_id" gorm:"primaryKey"`
	Student_name     string `json:"student_name" binding:"required" gorm:"not null"`
	Student_age      int    `json:"student_age" binding:"required" gorm:"not null"`
	Student_address  string `json:"student_address" binding:"required" gorm:"not null"`
	Student_phone_no string `json:"student_phone_no" binding:"required" gorm:"not null"`
}
