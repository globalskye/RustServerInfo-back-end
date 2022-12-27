package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetVkPosts(c *gin.Context) {
	posts, err := h.services.VkI.GetVkPosts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
	return
}
