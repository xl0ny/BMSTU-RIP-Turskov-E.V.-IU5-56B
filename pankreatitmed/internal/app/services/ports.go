package services

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"time"
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
	GetImageName(critid uint) (string, error)
}

type MedOrdersRepoPort interface {
	CountItems(orderID uint) (int64, error)
	IsMedOrderDeleted(orderID uint) (bool, error)
	IsMedOrderDraft(orderID uint) (bool, error)
	IsMedOrderFormed(orderID uint) (bool, error)

	GetOrCreateDraftMedOrder(creatorID uint) (*ds.MedOrder, error)
	GetMedOrders(status string, start, end time.Time) ([]ds.MedOrder, error)
	GetMedOrderWithItems(orderID uint) (ds.MedOrder, []ds.MedOrderItem, error)

	UpdateMedOrder(id uint, order *request.UpdateMedOrder) error
	FormMedOrder(id uint) error
	EndOrCancelMedOrder(id, moderator uint, status string) error
	SoftDeleteOrderSQL(orderID uint) error
	SetRansonAndRisk(orderID uint, score int, risk string) error
	GetCriterionByID(id uint) (*ds.Criterion, error)
}

type MedOrderItemsRepoPort interface {
	DeleteFromOrder(medorder, criterion uint) error
	UpdateMedOrderItem(medorder, criterion uint, position *uint, val *float64) error
}

type MedUsersRepoPort interface {
	CreateMedUser(user *ds.MedUser) error
	GetMedUserByLogin(login string) (*ds.MedUser, error)
	ChangeMedUser(id uint, user *request.UpdateMedUser) error
	GetMedUserByID(id uint) (*ds.MedUser, error)
}
