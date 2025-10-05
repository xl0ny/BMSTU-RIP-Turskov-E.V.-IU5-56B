package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const demoUserID uint = 1 // пока без авторизации

// GET /criteria?q=&order_id=
func (h *Handler) CriteriaList(c *gin.Context) {
	q := c.Query("q")
	orderIDStr := c.Query("med_order_id")
	var orderID uint
	if v, err := strconv.Atoi(orderIDStr); err == nil && v > 0 {
		orderID = uint(v)
	}
	checkdeleted, _ := h.Repository.IsMedOrderDeleted(orderID)
	if checkdeleted && orderID != 0 {
		c.HTML(http.StatusOK, "medordernotfind.html", gin.H{
			"Title": "Услуги (критерии Рэнсона)",
		})
		return
	}
	criteria, _ := h.Repository.GetCriteria(q)
	var cartCount int64
	if orderID > 0 {
		cartCount, _ = h.Repository.CountItems(orderID)
	}
	fmt.Print("ВВУВувцу")
	fmt.Print(cartCount)
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
	idStr := c.Query("id")
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
