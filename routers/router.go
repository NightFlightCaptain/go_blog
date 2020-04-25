package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go_blog/pkg/setting"
	"go_blog/routers/api"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.Config.Server.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	r.GET("/wait", func(context *gin.Context) {
		time.Sleep(10 * time.Second)
	})
	//r.Use(jwt.JWT())

	tags := r.Group("/tag")
	{
		tags.GET("", api.GetTags)
		tags.POST("", api.AddTag)
		tags.PUT("", api.EditTag)
		tags.DELETE("", api.DeleteTag)
	}

	articles := r.Group("/article")
	{
		articles.GET("/:id", api.GetArticle)
		articles.GET("/", api.GetArticles)
		articles.POST("/", api.AddArticle)
		articles.PUT("/:id", api.EditArticle)
		articles.DELETE("/:id", api.DeleteArticle)
	}

	return r
}
