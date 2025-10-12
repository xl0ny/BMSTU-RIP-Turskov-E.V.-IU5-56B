package handler

import (
	"net/http"
	"pankreatitmed/internal/app/dto/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteOrderItem(c *gin.Context) {
	var item request.GetMedOrderItem
	if err := c.ShouldBindQuery(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.svcs.MedOrderItems.Delete(item.MedOrderID, item.CriterionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdateOrderItem(c *gin.Context) {
	var item request.GetMedOrderItem
	var fields request.MedOrderItemUpdate
	if err := c.ShouldBindQuery(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedOrderItems.Update(item.MedOrderID, item.CriterionID, fields.Position, fields.ValueNum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
