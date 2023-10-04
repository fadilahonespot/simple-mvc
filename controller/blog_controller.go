package controller

import (
	"net/http"
	"simple-mvc/model"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type BlogController struct {
	DB *gorm.DB
}

func SetupBlogControler(db *gorm.DB) BlogController {
	return BlogController{db}
}

// get all blogs
func (s *BlogController) GetBlogsController(c echo.Context) error {
	var blogs []model.Blog

	if err := s.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// get blog by id
func (s *BlogController) GetBlogController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "blog id not valid",
		})
	}
	var blog model.Blog
	err = s.DB.Take(&blog, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "blog not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog by id",
		"blogs":   blog,
	})
}

// create new blog
func (s *BlogController) CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	c.Bind(&blog)

	var user model.User
	err := s.DB.Take(&user, "id = ?", blog.UserId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user id not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	if err := s.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

// delete blog by id
func (s *BlogController) DeleteBlogController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "blog id not valid",
		})
	}
	var blog model.Blog
	err = s.DB.Take(&blog, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "blog not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	err = s.DB.Delete(&blog, "id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

// update blog by id
func (s *BlogController) UpdateBlogController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "blog id not valid",
		})
	}

	var req model.Blog
	c.Bind(&req)

	var blog model.Blog
	err = s.DB.Take(&blog, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "blog not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	var user model.User
	err = s.DB.Take(&user, "id = ?", blog.UserId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "user id not found",
			})
		}
    return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "message": "internal server error",
    })
	}


	req.ID = uint(id)
	req.CreatedAt = blog.CreatedAt
	err = s.DB.Save(&req).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed update blog",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blogs":   req,
	})
}
