package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/middleware/jwt"
	"go_blog/pkg/setting"
	"go_blog/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)

	tags := r.Group("/tag").Use(jwt.JWT())
	{
		tags.GET("", api.GetTags)
		tags.POST("", api.AddTag)
		tags.PUT("", api.EditTag)
		tags.DELETE("", api.DeleteTag)
	}

	articles := r.Group("/article").Use(jwt.JWT())
	{
		articles.GET("/:id", api.GetArticle)
		articles.GET("/", api.GetArticles)
		articles.POST("/", api.AddArticle)
		articles.PUT("/:id", api.EditArticle)
		articles.DELETE("/:id", api.DeleteArticle)
	}

	return r
}
