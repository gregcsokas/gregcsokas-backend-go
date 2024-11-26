package newsletter

import "github.com/gin-gonic/gin"

type Handler struct {
	module *Module
}

func NewHandler(module *Module) *Handler {
	return &Handler{
		module: module,
	}
}

func (h *Handler) Subscribe(context *gin.Context)           {}
func (h *Handler) UnSubscribe(context *gin.Context)         {}
func (h *Handler) GetSubscriptionInfo(context *gin.Context) {}
