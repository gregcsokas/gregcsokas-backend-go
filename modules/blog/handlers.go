package blog

import "github.com/gin-gonic/gin"

type Handler struct {
	module *Module
}

func NewHandler(module *Module) *Handler {
	return &Handler{
		module: module,
	}
}

func (h *Handler) GetCategories(context *gin.Context) {}
func (h *Handler) GetTags(context *gin.Context)       {}
func (h *Handler) CreatePost(context *gin.Context)    {}
func (h *Handler) GetPosts(context *gin.Context)      {}
func (h *Handler) GetPost(context *gin.Context)       {}
func (h *Handler) UpdatePost(context *gin.Context)    {}
