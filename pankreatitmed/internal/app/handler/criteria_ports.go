package handler

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
)

type CriteriaRepoPort interface {
	GetCriteria(q string) ([]ds.Criterion, error)
	GetCriterionByID(id uint) (*ds.Criterion, error)
	CreateCriterion(c *ds.Criterion) error
	UpdateCriterion(id uint, in *request.UpdateCriterion) error
	DeleteCriterion(id uint) error
	AddItem(orderID, criterionID uint) error
}
