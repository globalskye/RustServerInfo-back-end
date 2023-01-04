package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllPlayers(c *gin.Context) {
	players, err := h.services.PlayerI.GetPlayers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
	return
}
func (h *Handler) GetPlayerBySteamId(c *gin.Context) {
	steamId := c.Param("steamId")
	id, err := strconv.Atoi(steamId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Player, err := h.services.PlayerI.GetPlayerBySteamId(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Player)
	return
}
func (h *Handler) GetPlayerByName(c *gin.Context) {
	name := c.Param("name")

	player, err := h.services.PlayerI.GetPlayerByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, player)
	return
}

func (h *Handler) GetOnline(c *gin.Context) {
	online, err := h.services.PlayerI.GetOnline()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, online)
	return

}
func (h *Handler) GetTopRaiders(c *gin.Context) {
	players, err := h.services.PlayerI.GetTopRaiders()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
	return
}

func (h *Handler) GetTopKillers(c *gin.Context) {
	players, err := h.services.PlayerI.GetTopKillers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
	return
}

func (h *Handler) GetTopOnline(c *gin.Context) {

	players, err := h.services.PlayerI.GetTopTime()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
	return
}
