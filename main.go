package main

import (
	"simple-mvc/config"
	"simple-mvc/controller"
	"simple-mvc/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	_ "github.com/joho/godotenv"
)

func main() {
	db := config.InitDB()
	userController := controller.SetupUserControler(db)
	blogController := controller.SetupBlogControler(db)

	// create a new echo instance
	e := echo.New()
	router.NewRouter(e, userController, blogController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
