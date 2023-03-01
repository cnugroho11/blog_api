package controller

import (
	"net/http"

	"github.com/cnugroho11/blog_api/model"
	"github.com/cnugroho11/blog_api/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BlogController struct {
	DB *gorm.DB
}

func NewBlogController(DB *gorm.DB) BlogController {
	return BlogController{DB}
}

func (bc *BlogController) GetBlog(ctx echo.Context) error {
	var blogs []model.Blog

	getBlogs := bc.DB.Find(&blogs)
	if getBlogs.Error != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "failed",
			Data:    &echo.Map{},
		})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"blog": blogs,
		},
	})
}

func (bc *BlogController) GetBySlug(ctx echo.Context) error {
	var blog model.Blog
	var slugParam string = ctx.Param("slug")

	getBlog := bc.DB.Where("slug = ?", slugParam).Find(&blog)
	if getBlog.Error != nil {
		return ctx.JSON(http.StatusBadRequest, response.UserResponse{
			Status:  http.StatusBadRequest,
			Message: "failed",
			Data:    &echo.Map{},
		})
	}

	return ctx.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"blog": blog,
		},
	})
}
