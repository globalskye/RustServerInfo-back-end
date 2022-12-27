package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func (h *Handler) GetUserBySteamId(c *gin.Context) {
	steamId := c.Param("steamId")
	id, err := strconv.Atoi(steamId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.services.UserI.GetUserBySteamId(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
func (h *Handler) GetUserByName(c *gin.Context) {
	name := c.Param("name")

	user, err := h.services.UserI.GetUserByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
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
func (h *Handler) GetTopRaiders(c *gin.Context) {
	users, err := h.services.UserI.GetTopRaiders()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

func (h *Handler) GetTopKillers(c *gin.Context) {
	users, err := h.services.UserI.GetTopKillers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}

func (h *Handler) GetTopOnline(c *gin.Context) {

	users, err := h.services.UserI.GetTopTime()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
