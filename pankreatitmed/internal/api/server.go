package api

import (
	"pankreatitmed/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	h.RegisterStatic(r)
	h.RegisterRoutes(r)
	return r
}
