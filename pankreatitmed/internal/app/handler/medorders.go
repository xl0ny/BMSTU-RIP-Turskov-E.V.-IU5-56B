package handler

import (
	"context"
	"net/http"
	"strconv"

	"pankreatitmed/internal/app/ds"

	"github.com/gin-gonic/gin"
)

// POST /order/add
func (h *Handler) MedOrderAdd(c *gin.Context) {
	criterionID, _ := strconv.Atoi(c.PostForm("criterion_id"))
	q := c.PostForm("q")
	o, _ := h.Repository.GetOrCreateDraftMedOrder(demoUserID)
	medorderID := int(o.ID)
	_ = h.Repository.AddItem(uint(medorderID), uint(criterionID))
	c.Redirect(http.StatusFound, "/criteria?q="+q)
}

// GET /order?id=
func (h *Handler) MedOrderView(c *gin.Context) {
	idStr := c.Param("id")
	oid, _ := strconv.Atoi(idStr)

	o, items, err := h.Repository.GetMedOrderWithItems(uint(oid))
	if err != nil || o.Status == "deleted" {
		c.Redirect(http.StatusFound, "/criteria")
		return
	}

	critMap := map[uint]ds.Criterion{}
	for _, it := range items {
		if _, ok := critMap[it.CriterionID]; !ok {
			if cr, _ := h.Repository.GetCriterionByID(it.CriterionID); cr != nil {
				critMap[it.CriterionID] = *cr
			}
		}
	}

	r := len(items)
	o.RansonScore = &r
	c.HTML(http.StatusOK, "medorder.html", gin.H{
		"Title":       "Заявка",
		"MedOrderID":  o.ID,
		"MedOrder":    o,
		"CriteriaMap": critMap,
		"Items":       items,
		"Length":      r,
	})
}

// POST /order/delete  — SQL UPDATE (без ORM)
func (h *Handler) MedOrderDelete(c *gin.Context) {
	idStr := c.PostForm("med_order_id")
	oid, _ := strconv.Atoi(idStr)
	_ = h.Repository.SoftDeleteOrderSQL(context.Background(), uint(oid))
	c.Redirect(http.StatusFound, "/criteria")
}
