package router

import (
	"simple-mvc/controller"

	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, userController controller.UserController, blogController controller.BlogController) {
	// Route / to handler function
	e.GET("/users", userController.GetUsersController)
	e.GET("/users/:id", userController.GetUserController)
	e.POST("/users", userController.CreateUserController)
	e.DELETE("/users/:id", userController.DeleteUserController)
	e.PUT("/users/:id", userController.UpdateUserController)

	e.GET("/blogs", blogController.GetBlogsController)
	e.GET("/blogs/:id", blogController.GetBlogController)
	e.POST("/blogs", blogController.CreateBlogController)
	e.DELETE("/blogs/:id", blogController.DeleteBlogController)
	e.PUT("/blogs/:id", blogController.UpdateBlogController)
}