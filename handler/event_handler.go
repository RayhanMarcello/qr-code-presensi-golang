package handler

import (
	"presensi-qr-backend/service"

	"github.com/gin-gonic/gin"
)

type handlerEvent struct {
	srv service.EventService
}

func NewHandlerEvent(srv service.EventService) *handlerEvent {
	return &handlerEvent{srv: srv}
}

type reqEvent struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
	CreatedBy   uint   `json:"created_by"`
}

func (h *handlerEvent) CreateEventHandler(c *gin.Context) {
	var req reqEvent
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "isi semua field brow",
		})
	}
	result, err := h.srv.CreateEvent(req.Title, req.Description, req.StartAt, req.EndAt, req.CreatedBy)
	if err != nil {
		c.JSON(500, gin.H{
			"err": "cannot send to server",
		})
	}

	c.JSON(200, gin.H{
		"message": "sucsess",
		"data":    result,
	})

}
