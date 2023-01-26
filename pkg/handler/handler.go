package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalskye/RustServerInfo-back-end.git/pkg/service"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	gin.SetMode(gin.ReleaseMode)

	gin.Logger()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/userExist/:name", h.checkUserName)
	}

	api := router.Group("/api")
	{
		api.GET("/players", h.GetAllPlayers)
		api.GET("/clans", h.GetAllClans)
		api.GET("/online", h.GetOnline)
		api.GET("/topKillers", h.GetTopKillers)
		api.GET("/topClans", h.GetTopClans)
		api.GET("/topRaiders", h.GetTopRaiders)
		api.GET("/topOnline", h.GetTopOnline)
		api.GET("/playerBySteamId/:steamId", h.GetPlayerBySteamId)
		api.GET("/playerByName/:name", h.GetPlayerByName)
		api.GET("/clanByName/:name", h.GetClanByName)
		api.GET("/vk", h.GetVkPosts)
		api.GET("/client", h.GameClientDownload)

		shop := api.Group("/shop")
		{
			shop.GET("/all", h.GetALlShopItems)
			shop.GET("/:category", h.GetShopItemsByCategory)
			shop.POST("/", h.InsertItemIntoShop)
		}
	}
	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/", h.GetUser)
		user.GET("/name", h.GetUserName)

	}

	return router

}
