package http

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type getAllOrdersResponse struct {
	Data []model.Order `json:"data"`
}

// @Summary Get All
// @Tags order
// @Description get all orders
// @ID getall-orders
// @Accept json
// @Produce json
// @Success 200 {object} model.Order
// @Failure default {object} errorResponse
// @Router /api/order [get]
func (h *Handler) getAllOrders(c *gin.Context) {
	orders := h.services.GetAllFromCache()
	c.JSON(http.StatusOK, getAllOrdersResponse{Data: orders})
}

// @Summary Get By Uid
// @Tags order
// @Description Get order by uid
// @ID getbyuid-order
// @Accept json
// @Produce json
// @Param uid path string true "Order uid"
// @Success 200 {object} model.Order
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/order/{uid} [get]
func (h *Handler) getOrderByUid(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.services.GetFromCache(uid)
	if err != nil {
		if val, ok := err.(cache.ErrorHandler); ok {
			newErrorResponse(c, val.StatusCode, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary Get By Uid From DB
// @Tags DB
// @Description Get order by uid from DB
// @ID getbyuid-db-order
// @Accept json
// @Produce json
// @Param uid path string true "Order uid"
// @Success 200 {object} model.Order
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/order/db/{uid} [get]
func (h *Handler) getOrderByUidFromDB(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.services.GetFromDB(uid)
	if err != nil {
		log.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}
