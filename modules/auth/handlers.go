package auth

import "github.com/gin-gonic/gin"

type Handler struct {
	module *Module
}

func NewHandler(module *Module) *Handler {
	return &Handler{
		module: module,
	}
}

func (h *Handler) Register(context *gin.Context) {}

func (h *Handler) Login(context *gin.Context) {}

func (h *Handler) RefreshToken(context *gin.Context) {}

func (h *Handler) Logout(context *gin.Context) {}

func (h *Handler) ForgotPassword(context *gin.Context) {}

func (h *Handler) ResetPassword(context *gin.Context) {}
