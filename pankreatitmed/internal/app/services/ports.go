package services

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
	GetSeq() (uint, error)
	ResetCriterionSequence() error
	GetOrCreateDraftMedOrder(creatorID uint) (*ds.MedOrder, error)
}

// При необходимости добавишь сюда порты заказов/пользователей и т.д.
// type OrdersRepoPort interface { ... }
// type UsersRepoPort  interface { ... }

// Контекст обычно протаскиваем из хендлера, но раз твои репо не требуют ctx — можно опустить.
//_ = context.Background
