package http

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type getAllOrdersResponse struct {
	Data []model.Order `json:"data"`
}

func (h *Handler) getAllOrders(c *gin.Context) {
	orders := h.services.GetAllFromCache()
	c.JSON(http.StatusOK, getAllOrdersResponse{Data: orders})
}

func (h *Handler) getOrderByUid(c *gin.Context) {
	uid := c.Param("uid")

	order, err := h.services.GetFromCache(uid)
	if err != nil {
		log.Print(err)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

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
