package handler

import (
	"fmt"
	"net/http"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/singleton"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PankreatitOrderFromCart(c *gin.Context) {
	mo, err := h.svcs.PankreatitOrders.GetDraft(singleton.GetCurrentUser().ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mo)
}

func (h *Handler) ListPankreatitOrders(c *gin.Context) {
	var filters request.GetPankreatitOrders
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("--------------------")
	fmt.Println(filters.Status, filters.FromDate, filters.ToDate)
	res, err := h.svcs.PankreatitOrders.List(filters.Status, filters.FromDate, filters.ToDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) PankreatitOrderGet(c *gin.Context) {
	var id request.GetPankreatitOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svcs.PankreatitOrders.Get(id.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

// TODO разобраться почему не выводит ошибку, когда не находит ордер по id-ку
func (h *Handler) PankreatitOrderUpdate(c *gin.Context) {
	var id request.GetPankreatitOrder
	var mo request.UpdatePankreatitOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&mo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(mo.Status)
	if err := h.svcs.PankreatitOrders.Update(id.ID, &mo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) PankreatitOrderForm(c *gin.Context) {
	var id request.GetPankreatitOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.svcs.PankreatitOrders.Form(id.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) PankreatitOrderComplete(c *gin.Context) {
	var idstatus request.EndOrCancelPankreatitOrder
	var moderator request.GetModerator
	if err := c.ShouldBindUri(&idstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&moderator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.PankreatitOrders.CancelOrEnd(idstatus.ID, moderator.ModeratorID, moderator.Password, idstatus.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) PankreatitOrderDelete(c *gin.Context) {
	var id request.GetPankreatitOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.PankreatitOrders.Delete(id.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
