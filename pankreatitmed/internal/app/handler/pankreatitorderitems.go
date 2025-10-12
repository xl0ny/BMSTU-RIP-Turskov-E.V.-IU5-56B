package handler

import (
	"net/http"
	"pankreatitmed/internal/app/dto/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeletePankreatitOrderItem(c *gin.Context) {
	var item request.GetPankreatitOrderItem
	if err := c.ShouldBindQuery(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.svcs.PankreatitOrderItems.Delete(item.PankreatitOrderID, item.CriterionID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdatePankreatitOrderItem(c *gin.Context) {
	var item request.GetPankreatitOrderItem
	var fields request.PankreatitOrderItemUpdate
	if err := c.ShouldBindQuery(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.PankreatitOrderItems.Update(item.PankreatitOrderID, item.CriterionID, fields.Position, fields.ValueNum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
