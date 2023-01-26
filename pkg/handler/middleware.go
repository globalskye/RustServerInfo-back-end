package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

func getUserId(c *gin.Context) (primitive.ObjectID, error) {
	id, ok := c.Get("userId")
	if !ok {
		return primitive.ObjectID{}, errors.New("user id not found")
	}

	userId, ok := id.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, errors.New("user id is of invalid type")
	}

	return userId, nil
}
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "auth header left")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
		return
	}
	if headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token type, should be bearer")
		return
	}
	userId, err := h.services.Authorization.ParseAccessToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}
	user, err := h.services.UserI.GetUserById(userId)
	if err != mongo.ErrNoDocuments {
		c.Set("role", user.Role)
		c.Set("userId", userId)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errors.New("Token exists, but user in database not found").Error())
		return
	}

	return
}
