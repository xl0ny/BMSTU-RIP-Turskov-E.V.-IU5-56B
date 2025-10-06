package request

import (
	"time"
)

type GetMedOrders struct {
	Status   string    `json:"status" binding:"required, oneof=formed completed rejected"`
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
}

type GetMedOrder struct {
	ID uint `json:"id" binding:"required"`
}

type UpdateMedOrder struct {
	Status        string    `json:"status" binding:"omitempty, oneof=draft deleted formed completed rejected"`
	RansonScore   int       `json:"ranson_score"`
	MortalityRisk string    `json:"mortality_risk"`
}

type FormMedOrder struct {
	ID uint `json:"id" binding:"required"`
	//CreatorID uint   `json:"creator_id" binding:"required"`
	//Password  string `json:"password" binding:"required"`
}

type EndOrCancelMedOrder struct {
	ID          uint   `json:"id" binding:"required"`
	ModeratorID uint   `json:"moderator_id" binding:"required"`
	Status      string `json:"status" binding:"required, oneof=completed rejected"`
}

type DeleteMedOrder struct {
	ID uint `json:"id" binding:"required"`
}
