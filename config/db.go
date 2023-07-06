package config

import (
	"fmt"

	"github.com/Hasyim-Kai/Go_Gin_Rest_Api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbInit() *gorm.DB {
	dsn := "host=localhost user=postgres password=123 dbname=practice_rest_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal konek ke db")
	}

	fmt.Println("Sukses Konek ke Db!")
	DB.AutoMigrate(&model.Student{})
	data := model.Student{}
	if err := DB.Limit(1).Find(&data).Error; err != nil {
		fmt.Println(err)
		fmt.Println("=================== run seeder user ======================")
		seederUser()
	}
	return DB
}

func seederUser() {
	data := model.Student{
		// Student_id:       1,
		Student_name:     "Dono",
		Student_age:      20,
		Student_address:  "Jakarta",
		Student_phone_no: "0123456789",
	}
	DB.Create(&data)
}
