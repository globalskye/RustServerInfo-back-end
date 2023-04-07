package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"os"
	"time"
)

type AuthService struct {
	repo repository.UserI
}

func NewAuthService(repo repository.UserI) *AuthService {
	return &AuthService{repo: repo}
}

type jwtTokenClaims struct {
	jwt.StandardClaims
	UserId primitive.ObjectID `json:"_id"`
	Role   string             `json:"role"`
}

func generateHashPassword(password string) string {
	
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}

func (a *AuthService) GenerateAccessToken(username, password string) (string, error) {
	user, err := a.repo.GetUserByCredentials(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwtTokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(1680000 * time.Hour).Unix(), // 7days
		},
		UserId: user.Id,
		Role:   user.Role,
	})
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
func (a *AuthService) ParseAccessToken(accessToken string) (primitive.ObjectID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return primitive.ObjectID{}, err
	}
	claims, ok := token.Claims.(*jwtTokenClaims)
	if !ok {
		return primitive.ObjectID{}, errors.New("bad token claims type")
	}

	return claims.UserId, nil

}
