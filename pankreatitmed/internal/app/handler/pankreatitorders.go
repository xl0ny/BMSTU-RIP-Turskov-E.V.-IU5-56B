package handler

import (
	"net/http"
	"pankreatitmed/internal/app/authctx"
	"pankreatitmed/internal/app/dto/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PankreatitOrderFromCart(c *gin.Context) {
	usr, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	mo, err := h.svcs.PankreatitOrders.GetDraft(usr.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, mo)
}

func (h *Handler) ListPankreatitOrders(c *gin.Context) {
	usr, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	var filters request.GetPankreatitOrders
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svcs.PankreatitOrders.List(usr.ID, filters.Status, filters.FromDate, filters.ToDate)
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
	if err := h.svcs.PankreatitOrders.Update(id.ID, &mo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) PankreatitOrderForm(c *gin.Context) {
	usr, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	var id request.GetPankreatitOrder
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	o, err := h.svcs.PankreatitOrders.Get(id.ID)
	//fmt.Println(o.CreatorID, usr.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if o.CreatorID != usr.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "u don't have permission to form this order(not your order)"})
		return
	}
	if err := h.svcs.PankreatitOrders.Form(id.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) PankreatitOrderComplete(c *gin.Context) {
	var idstatus request.EndOrCancelPankreatitOrder
	usr, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	if err := c.ShouldBindUri(&idstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.PankreatitOrders.CancelOrEnd(idstatus.ID, usr.ID, idstatus.Status); err != nil {
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
