package rest

import (
	"farm/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	g *gin.Engine
	R *services.Refrigerator
}

func NewHandler() *Handler {
	return &Handler{g: gin.Default()}
}

func (h *Handler) Start() {
	h.g.GET("/refrigerator/eggs", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Eggs": h.R.Eggs})
	})
	h.g.Run("127.0.0.1:8000")
}
