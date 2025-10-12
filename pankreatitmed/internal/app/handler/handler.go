package handler

import (
	"pankreatitmed/internal/app/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct{ Repository *repository.Repository }

func NewHandler(r *repository.Repository) *Handler { return &Handler{Repository: r} }

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/criteria", h.CriteriaList)         // список
	r.GET("/criterion/:id", h.CriterionDetail) // деталка
	r.GET("/medorder/:id", h.MedOrderView)     // заявка

	r.POST("/medorder/add", h.MedOrderAdd)       // ORM
	r.POST("/medorder/delete", h.MedOrderDelete) // SQL UPDATE
}

func (h *Handler) RegisterStatic(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static/styles", "./resources/styles")
}
