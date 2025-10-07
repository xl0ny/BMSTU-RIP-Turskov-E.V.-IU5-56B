package handler

import (
	"net/http"
	"pankreatitmed/internal/app/dto"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/mapper"

	"github.com/gin-gonic/gin"
)

//const demoUserID uint = 1 // пока без авторизации

// CriteriaList GET /criteria
func (h *Handler) CriteriaList(c *gin.Context) {
	var query request.GetCriteria
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list, err := h.svcs.Criteria.List(query.Query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	items := mapper.CriterionsToSendCrtierions(list)
	res := dto.List[response.SendCriterion]{Items: items}

	c.JSON(http.StatusOK, res)

}

// CriteriaGet GET /criteria
func (h *Handler) CriteriaGet(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println("id", id.ID)
	criterion, err := h.svcs.Criteria.Get(id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	crit := mapper.CritertionToSendCriterionLink(criterion)
	c.JSON(http.StatusOK, crit)
}

func (h *Handler) CriteriaCreate(c *gin.Context) {
	var criterion request.CreateCriterion
	if err := c.ShouldBindJSON(&criterion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		println(1)
		return
	}
	crit, err := mapper.CreateCriterionToCriterion(criterion)
	if err != nil {
		println(2)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if err := h.svcs.Criteria.Create(&crit); err != nil {
		println(3)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

// CriteriaUpdate TODO разобраться почему не кидает ошибку
func (h *Handler) CriteriaUpdate(c *gin.Context) {
	var id request.GetCriterion
	var criterion request.UpdateCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&criterion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.Update(id.ID, &criterion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (h *Handler) CriteriaDelete(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.Delete(id.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

// TODO настроить нормализацию последовательности БД
func (h *Handler) AddCriteriaToDraft(c *gin.Context) {
	var id request.GetCriterion
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.Criteria.ToDradt(id.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
