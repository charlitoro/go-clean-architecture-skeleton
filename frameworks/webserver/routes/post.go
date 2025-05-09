package routes

import (
	"github.com/charlitoro/go-clean-architecture-skeleton/adapters/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(reoute *gin.Engine) {
	postController := controllers.NewPostController()

	reoute.POST("/api/v1/post", postController.AddPost)
}
