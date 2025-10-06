package handler

import (
	"net/http"
	"pankreatitmed/internal/app/dto"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/mapper"

	"github.com/gin-gonic/gin"
)

const demoUserID uint = 1 // пока без авторизации

func (h *Handler) CriteriaList(c *gin.Context) {
	var query request.GetCriteria
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//list, err := h.svc
	list, err := h.Repository.GetCriteria(query.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	items := mapper.CriterionsToSendCrtierions(list)
	res := dto.List[response.SendCriterion]{Items: items}

	c.JSON(http.StatusOK, res)

}

// GET /criteria
