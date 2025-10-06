// internal/app/service/criteria_service.go
package services

import (
	"context"
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
)

// минимальный контракт репозитория, НУЖНЫЙ именно этому сервису
type CriteriaRepo interface {
	GetCriteria(q string) ([]ds.Criterion, error)
	GetCriterionByID(id uint) (*ds.Criterion, error)
	CreateCriterion(c *ds.Criterion) error
	UpdateCriterion(id uint, in *request.UpdateCriterion) error
	DeleteCriterion(id uint) error
	AddItem(orderID, criterionID uint) error
}

type CriteriaServiceImpl struct {
	repo CriteriaRepo
}

func NewCriteriaService(repo CriteriaRepo) *CriteriaServiceImpl {
	return &CriteriaServiceImpl{repo: repo}
}

func (s *CriteriaServiceImpl) List(ctx context.Context, q string) ([]ds.Criterion, error) {
	// тут можно навесить бизнес-правила/валидацию/кэш и т.д.
	return s.repo.GetCriteria(q)
}
