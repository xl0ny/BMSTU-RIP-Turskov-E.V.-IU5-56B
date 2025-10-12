package handler

import (
	"pankreatitmed/internal/app/services"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svcs *services.Services }

func NewHandler(svcs *services.Services) *Handler { return &Handler{svcs: svcs} }

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		srv := api.Group("/criteria")
		{
			srv.GET("", h.CriteriaList)
			srv.GET("/:id", h.CriteriaGet)
			srv.POST("", h.CriteriaCreate)
			srv.PUT("/:id", h.CriteriaUpdate)
			srv.DELETE("/:id", h.CriteriaDelete)
			srv.POST("/:id/image", h.UploadCriterionImage)
			srv.POST("/:id/add-to-draft", h.AddCriteriaToDraft)
		}

		ord := api.Group("/pankreatitorders")
		{
			ord.GET("/cart", h.PankreatitOrderFromCart)
			ord.GET("", h.ListPankreatitOrders)
			ord.GET(":id", h.PankreatitOrderGet)
			ord.PUT("/:id", h.PankreatitOrderUpdate)
			ord.PUT("/:id/form", h.PankreatitOrderForm)
			ord.PUT("/:id/set/:status", h.PankreatitOrderComplete)
			ord.DELETE("/:id", h.PankreatitOrderDelete)

			ord.DELETE("/items", h.DeletePankreatitOrderItem)
			ord.PUT("/items", h.UpdatePankreatitOrderItem)
		}

		auth := api.Group("/users")
		{
			auth.POST("auth/register", h.MedUserRegistation)
			auth.GET("me", h.MedUserGetFields)
			auth.PUT("me", h.MedUserUpdateFields)
			auth.POST("auth/login", h.MedUserLogIn)
			auth.POST("auth/logout", h.MedUserLogOut)
		}
	}
}

//func (h *Handler) RegisterStatic(r *gin.Engine) {
//	r.LoadHTMLGlob("templates/*")
//	r.Static("/static/styles", "./resources/styles")
//}
