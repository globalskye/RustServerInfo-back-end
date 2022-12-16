package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.services.UserI.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
func (h *Handler) GetOnline(c *gin.Context) {
	online, err := h.services.UserI.GetOnline()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, online)
	return

}
