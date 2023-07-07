package config

import (
	"fmt"

	"github.com/Hasyim-Kai/Go_Gin_Rest_Api_MySkill/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=123 dbname=practice_rest_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal konek ke db")
	}
	fmt.Println("Sukses Konek ke Db!")

	DB.AutoMigrate(&model.Student{})

	var Student []model.Student
	DB.Limit(1).Find(&Student)

	if len(Student) == 0 {
		fmt.Println("=================== Run Seeder Student ======================")
		seed := []model.Student{
			{
				Student_name:     "Dono",
				Student_age:      20,
				Student_address:  "Jakarta",
				Student_phone_no: "0123456789",
			},
			{
				Student_name:     "Joko",
				Student_age:      29,
				Student_address:  "Jonggol",
				Student_phone_no: "0123456789",
			}}

		if res := DB.Create(seed); res.Error == nil {
			fmt.Println("=================== Seeder Success ======================")
		}
	}
}
