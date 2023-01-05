package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/model"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

func (a *AuthService) CheckUserName(name string) (bool, error) {
	user, err := a.repo.GetUserByName(name)
	fmt.Println(user)
	fmt.Println(err)
	return false, err
}

func (a *AuthService) GetUserById(id int) ([]model.User, error) {
	return a.repo.GetUserById(id)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type jwtTokenClaims struct {
	jwt.StandardClaims
	UserId primitive.ObjectID `json:"_id"`
}

func (a *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	user.Id = primitive.NewObjectID()
	return a.repo.CreateUser(user)

}
func generateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}

func (a *AuthService) GenerateAccessToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwtTokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(168 * time.Hour).Unix(), // 7days
		},
		UserId: user.Id,
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
