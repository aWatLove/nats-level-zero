package http

import (
	"github.com/aWatLove/nats-lvl-zero/internal/service"
	"github.com/gin-gonic/gin"
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

	return router

}
