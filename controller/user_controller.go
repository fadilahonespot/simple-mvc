package controller

import (
	"net/http"
	"simple-mvc/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type UserController struct {
	DB *gorm.DB
}

func SetupUserControler(db *gorm.DB) UserController {
	return UserController{db}
}

// get all users
func (s *UserController) GetUsersController(c echo.Context) error {
	var users []model.User

	if err := s.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func (s *UserController) GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user id not valid",
		})
	}
	var user model.User
	err = s.DB.Take(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "internal server error",
    })
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"users":   user,
	})
}

// create new user
func (s *UserController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := s.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func (s *UserController) DeleteUserController(c echo.Context) error {
	// your solution here
  id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user id not valid",
		})
	}
	var user model.User
	err = s.DB.Take(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "internal server error",
    })
	}

  err = s.DB.Delete(&user, "id = ?", id).Error
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "internal server error",
    })
  }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func (s *UserController) UpdateUserController(c echo.Context) error {
	// your solution here
  id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "user id not valid",
		})
	}

  var req model.User
  c.Bind(&req)

	var user model.User
	err = s.DB.Take(&user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user not found",
			})
		}
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "internal server error",
    })
	}

  req.ID = uint(id)
  req.CreatedAt = user.CreatedAt
  err = s.DB.Save(&req).Error
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]interface{}{
      "message": "failed update user",
    })
  }
	
  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
    "users": req,
	})
}