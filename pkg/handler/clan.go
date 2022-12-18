package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllClans(c *gin.Context) {
	clans, err := h.services.ClanI.GetClans()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clans)
	return
}

func (h *Handler) GetTopClans(c *gin.Context) {
	clans, err := h.services.ClanI.GetTopClans()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clans)
	return
}
func (h *Handler) GetClanByName(c *gin.Context) {
	name := c.Param("name")
	clan, err := h.services.ClanI.GetClanByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clan)
	return
}
