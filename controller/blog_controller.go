package controller

import (
	"net/http"
	"time"

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
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  http.StatusBadRequest,
			Message: "failed",
			Data:    &echo.Map{},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"blog": blogs,
		},
	})
}

func (bc *BlogController) GetBySlug(ctx echo.Context) error {
	var slugParam string = ctx.Param("slug")
	var blog model.Blog

	getBlog := bc.DB.Where("slug = ?", slugParam).Find(&blog)
	if getBlog.Error != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  http.StatusBadRequest,
			Message: "failed",
			Data:    &echo.Map{},
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "success",
		Data: &echo.Map{
			"blog": blog,
		},
	})
}

func (bc *BlogController) CreateBlog(ctx echo.Context) error {
	var blog model.RequestBlog

	if err := ctx.Bind(&blog); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  http.StatusBadRequest,
			Message: "failed",
			Data:    &echo.Map{},
		})
	}

	newBlog := model.Blog{
		Title:     blog.Title,
		Slug:      blog.Slug,
		Content:   blog.Content,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createBlog := bc.DB.Create(&newBlog)
	if createBlog.Error != nil {
		return ctx.JSON(http.StatusBadRequest, response.Response{
			Status:  http.StatusBadRequest,
			Message: "failed insert data",
			Data:    &echo.Map{},
		})
	}

	return ctx.JSON(http.StatusAccepted, response.Response{
		Status:  http.StatusAccepted,
		Message: "success",
		Data:    &echo.Map{},
	})
}
