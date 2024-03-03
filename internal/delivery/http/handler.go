package http

import (
	_ "github.com/aWatLove/nats-lvl-zero/docs"
	"github.com/aWatLove/nats-lvl-zero/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/order/:uid", h.getOrderByUid)
		api.GET("/order/db/:uid", h.getOrderByUidFromDB)
		api.GET("/order", h.getAllOrders)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
	return router
}
