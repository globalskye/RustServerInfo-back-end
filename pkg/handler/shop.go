package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetALlShopItems(c *gin.Context) {
	items, err := h.services.ShopI.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
	return

}
