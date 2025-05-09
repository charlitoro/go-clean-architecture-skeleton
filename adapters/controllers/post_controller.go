package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/charlitoro/go-clean-architecture-skeleton/application/use_cases/post"
	"github.com/charlitoro/go-clean-architecture-skeleton/domain/entities"
)

type PostController struct {
	addPostUseCase *post.AddPostUseCase
}


func NewPostController() *PostController {
	addPostUseCase := post.NewAddPostUseCase()

	return &PostController{
		addPostUseCase: addPostUseCase,
	}
}

type AddPostRequest struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
    CreatedAt   string `json:"createdAt"` // optional, parse to time.Time
    IsPublished bool   `json:"isPublished"`
    UserId      string `json:"userId" binding:"required"`
}

func (c *PostController) AddPost(ctx *gin.Context) {
    var req AddPostRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Parse CreatedAt or use time.Now()
    createdAt := time.Now()
    if req.CreatedAt != "" {
        t, err := time.Parse(time.RFC3339, req.CreatedAt)
        if err == nil {
            createdAt = t
        }
    }

    input := entities.PostType{
        Title:       req.Title,
        Description: req.Description,
        CreatedAt:   createdAt,
        IsPublished: req.IsPublished,
        UserId:      req.UserId,
    }

    result := c.addPostUseCase.Execute(input)

    ctx.JSON(http.StatusOK, result)
}
