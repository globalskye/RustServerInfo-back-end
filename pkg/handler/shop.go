package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
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

func (h *Handler) GetShopItemsByCategory(c *gin.Context) {

}

func (h *Handler) InsertItemIntoShop(c *gin.Context) {
	var item model.DonatItem

	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.ShopI.InsertItem(item); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return
}
