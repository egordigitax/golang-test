package rest

import (
	"farm/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	g    *gin.Engine
	R    *services.Refrigerator
	port string
}

func NewHandler(port string) *Handler {
	return &Handler{g: gin.Default(), port: port}
}

func (h *Handler) Start() {
	h.g.GET("/refrigerator", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Eggs": h.R.Eggs})
	})
	h.g.Run("0.0.0.0:" + h.port)
}
