package config

import (
	"fmt"
	"os"
	"simple-mvc/model"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {

	config := model.Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{}, &model.Blog{})
	DB.Model(&model.Blog{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	
	DB.Debug()

	return DB
}