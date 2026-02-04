package handler

import (
	"presensi-qr-backend/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type Req struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req Req
	if err := c.ShouldBindBodyWithJSON(&req); err != nil || req.Name == "" || req.Email == "" || req.Password == "" {
		c.JSON(500, gin.H{
			"err": "error kontolodon",
		})
		return
	}
	u, err := h.svc.Create(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		if err == service.ErrEmailTaken {
			c.JSON(500, gin.H{
				"err": "email already exist brow",
			})
			return
		}
		c.JSON(404, gin.H{
			"message": "server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "sucsess create data brow",
		"data":    u,
	})
}
