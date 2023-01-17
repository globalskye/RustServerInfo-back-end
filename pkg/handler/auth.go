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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if exist := h.services.CheckUserName(user.Username); exist {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Пользователь уже существует"})
		return
	}
	id, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
	return
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
	return
}
func (h *Handler) checkUserName(c *gin.Context) {

	username := c.Param("name")

	if exist := h.services.CheckUserName(username); exist {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Пользователь уже существует"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Имя свободно"})
	return
}
