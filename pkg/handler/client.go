package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"net/http"
)

func (h *Handler) GameClientDownload(c *gin.Context) {
	var u model.User
	u.Cart = append(u.Cart, model.CartItem{})
	u.Cart[0].Item.Attachments = append(u.Cart[0].Item.Attachments, model.DonatAttachments{})

	u.Stock = append(u.Stock, model.StockItem{})
	u.Stock[0].Item.Attachments = append(u.Stock[0].Item.Attachments, model.DonatAttachments{})

	c.JSON(http.StatusOK, u)
}
