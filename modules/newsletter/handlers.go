package newsletter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Handler struct {
	module *Module
}

func NewHandler(module *Module) *Handler {
	return &Handler{
		module: module,
	}
}

func (h *Handler) Subscribe(context *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}

	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	userUuid, _ := uuid.NewV7()

	subscription := Subscription{
		Email: input.Email,
		UUID:  userUuid.String(),
	}

	err = h.module.db.Create(&subscription).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create subscription.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Subscription created.",
		"uuid":    subscription.UUID,
	})
}

func (h *Handler) UnSubscribe(context *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
	}

	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var subscription Subscription
	err = h.module.db.Where("email = ?", input.Email).Where("deleted_at IS NULL").First(&subscription).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Subscription not found.",
		})
		return
	}

	subscription.DeletedAt = time.Now()

	err = h.module.db.Save(&subscription).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update subscription.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Subscription deleted.",
	})
}

func (h *Handler) GetSubscriptionInfo(context *gin.Context) {
	userUuid := context.Param("uuid")

	fmt.Println(userUuid)

	var subscription Subscription
	err := h.module.db.Where("uuid = ?", userUuid).Where("deleted_at IS NULL").First(&subscription).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Subscription not found.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Subscription found.",
		"uuid":    subscription.UUID,
		"created": subscription.CreatedAt,
	})

}
