package handler

import (
	"fmt"
	"net/http"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/singleton"

	"github.com/gin-gonic/gin"
)

func (h *Handler) OrderFromCart(c *gin.Context) {
	mo, err := h.svcs.MedOrders.GetDraft(singleton.GetCurrentUser().ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mo)
}

func (h *Handler) ListOrders(c *gin.Context) {
	var filters request.GetMedOrders
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("--------------------")
	fmt.Println(filters.Status, filters.FromDate, filters.ToDate)
	res, err := h.svcs.MedOrders.List(filters.Status, filters.FromDate, filters.ToDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) OrderGet(c *gin.Context) {
	var id request.GetMedOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svcs.MedOrders.Get(id.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

// TODO разобраться почему не выводит ошибку, когда не находит ордер по id-ку
func (h *Handler) MedOrderUpdate(c *gin.Context) {
	var id request.GetMedOrder
	var mo request.UpdateMedOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&mo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(mo.Status)
	if err := h.svcs.MedOrders.Update(id.ID, &mo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) MedOrderForm(c *gin.Context) {
	var id request.GetMedOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svcs.MedOrders.Form(id.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) OrderComplete(c *gin.Context) {
	var idstatus request.EndOrCancelMedOrder
	var moderator request.GetModerator
	if err := c.ShouldBindUri(&idstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&moderator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedOrders.CancelOrEnd(idstatus.ID, moderator.ModeratorID, moderator.Password, idstatus.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) OrderDelete(c *gin.Context) {
	var id request.GetMedOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedOrders.Delete(id.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
