package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const demoUserID uint = 1 // пока без авторизации

// GET /criteria?q=&order_id=
func (h *Handler) CriteriaList(c *gin.Context) {
	q := c.Query("q")
	o, _ := h.Repository.GetOrCreateDraftMedOrder(demoUserID)
	orderID := uint(o.ID)
	checkdeleted, _ := h.Repository.IsMedOrderDeleted(orderID)
	if checkdeleted && orderID != 0 {
		c.HTML(http.StatusOK, "medordernotfind.html", gin.H{
			"Title": "Услуги (критерии Рэнсона)",
		})
		return
	}
	criteria, _ := h.Repository.GetCriteria(q)
	var cartCount int64
	cartCount, _ = h.Repository.CountItems(orderID)
	c.HTML(http.StatusOK, "criteria.html", gin.H{
		"Title":      "Услуги (критерии Рэнсона)",
		"Criteria":   criteria,
		"MedOrderID": orderID,
		"CartCount":  cartCount,
		"Query":      q,
	})
}

// GET /service?id=&order_id=
func (h *Handler) CriterionDetail(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.Atoi(idStr)
	orderIDStr := c.Query("med_order_id")
	orderID64, _ := strconv.Atoi(orderIDStr)

	cr, _ := h.Repository.GetCriterionByID(uint(id64))
	if cr == nil {
		c.String(http.StatusNotFound, "Критерий не найден")
		return
	}

	c.HTML(http.StatusOK, "criterion.html", gin.H{
		"Title":      cr.Name,
		"MedOrderID": uint(orderID64),
		"Criterion":  cr,
	})
}
