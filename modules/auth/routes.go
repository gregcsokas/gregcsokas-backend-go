package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, m *Module) {
	auth := router.Group("/auth")

	handler := NewHandler(m)

	// Public routes
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
	auth.POST("/refresh-token", handler.RefreshToken)

	// Password reset flow
	auth.POST("/forgot-password", handler.ForgotPassword)
	auth.POST("/reset-password", handler.ResetPassword)

	// Protected routes (require authentication)
	// authenticated := auth.Use(m.AuthMiddleware())
	auth.POST("/logout", handler.Logout)
}
