package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"net/http"
)

type userInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func (h *Handler) signIn(c *gin.Context) {
	var user userInput

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := h.services.Authorization.GenerateAccessToken(user.Username, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
func (h *Handler) checkUserName(c *gin.Context) {

}
