package handlers

import (
	"HexagonalModel/modules/project1/models"
	"HexagonalModel/modules/project1/services"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	PostRegisterHandler(c *gin.Context)
}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) PostRegisterHandler(c *gin.Context) {
	var req models.RequestRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err := h.s.PotsRegisterSer(req)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"status": "OK", "message": "User registered successfully"})
}
