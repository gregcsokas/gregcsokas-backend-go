package blog

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, m *Module) {
	blog := router.Group("/blog")

	handler := NewHandler(m)

	blog.POST("/", handler.CreatePost)
	blog.GET("/", handler.GetPosts)
	blog.GET("/:slug", handler.GetPost)
	blog.PUT("/:slug", handler.UpdatePost)
	blog.GET("/categories", handler.GetCategories)
	blog.GET("/tags", handler.GetTags)

}
