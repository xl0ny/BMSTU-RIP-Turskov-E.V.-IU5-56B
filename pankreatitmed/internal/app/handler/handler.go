package handler

import (
	"pankreatitmed/internal/app/services"

	"github.com/gin-gonic/gin"
)

type Handler struct{ svcs *services.Services }

func NewHandler(svcs *services.Services) *Handler { return &Handler{svcs: svcs} }

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	//r.GET("/criteria", h.CriteriaList)     // список
	//r.GET("/criterion", h.CriterionDetail) // деталка
	//r.GET("/medorder", h.MedOrderView)     // заявка
	//
	//r.POST("/medorder/add", h.MedOrderAdd)       // ORM
	//r.POST("/medorder/delete", h.MedOrderDelete) // SQL UPDATE
	api := r.Group("/api")
	{
		srv := api.Group("/criteria")
		{
			srv.GET("", h.CriteriaList)
			srv.GET("/:id", h.CriteriaGet)
			srv.POST("", h.CriteriaCreate)
			srv.PUT("/:id", h.CriteriaUpdate)
			srv.DELETE("/:id", h.CriteriaDelete)
			//srv.POST("/:id/image", h.Criteria.UploadImage)
			srv.POST("/:id/add-to-draft", h.AddCriteriaToDraft)
		}

		//ord := api.Group("/medorders")
		//{
		//	ord.GET("", h.Orders.CartIcon)
		//	ord.GET("", h.Orders.List)
		//	ord.GET("/:id", h.Orders.Get)
		//	ord.PUT("/:id", h.Orders.Update)
		//	ord.PUT("/:id/form", h.Orders.Form)
		//	ord.PUT("/:id/complete", h.Orders.Complete)
		//	ord.PUT("/:id/reject", h.Orders.Reject)
		//	ord.DELETE("/:id", h.Orders.Delete)
		//
		//	ord.PUT("/:id/items", h.OrderItems.Upsert)
		//	ord.DELETE("/:id/items", h.OrderItems.Delete)
		//}
		//
		//auth := api.Group("/users")
		//{
		//	auth.POST("auth/register", h.Users.Register)
		//	auth.POST("auth/login", h.Users.Login)
		//	auth.POST("auth/logout", h.Users.Logout)
		//	auth.GET("me", h.Users.Me)
		//	auth.PUT("me", h.Users.UpdateMe)
		//}
	}
}

func (h *Handler) RegisterStatic(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static/styles", "./resources/styles")
}
