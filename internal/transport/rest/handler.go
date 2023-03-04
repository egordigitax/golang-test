package rest

import (
	"farm/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	g *gin.Engine
	r *services.Refrigerator
}

func NewHandler() *Handler {
	return &Handler{g: gin.Default(), r: services.NewRefrigerator()}
}

func (h *Handler) Start() {
	h.g.GET("/refrigerator/eggs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	h.g.Run("127.0.0.1:8000")
}
