package config

import (
	"fmt"
	"simple-mvc/model"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {

	config := model.Config{
		DB_Username: "fadilah",
		DB_Password: "fadilah123",
		DB_Port:     "3306",
		DB_Host:     "db4free.net",
		DB_Name:     "fadilah_db",
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