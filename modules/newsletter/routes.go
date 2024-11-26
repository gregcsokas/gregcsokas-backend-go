package newsletter

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, m *Module) {
	newsletter := router.Group("/newsletter")

	handler := NewHandler(m)

	newsletter.POST("/subscribe", handler.Subscribe)
	newsletter.POST("/unsubscribe", handler.UnSubscribe)
	newsletter.GET("/:uuid", handler.GetSubscriptionInfo)
}
