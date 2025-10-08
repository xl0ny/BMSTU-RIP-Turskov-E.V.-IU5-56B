package request

import (
	"time"
)

type GetMedOrders struct {
	Status   string    `form:"status" binding:"required,oneof=formed completed rejected"`
	FromDate time.Time `form:"from_date" time_format:"2006-01-02T15:04:05"`
	ToDate   time.Time `form:"to_date" time_format:"2006-01-02T15:04:05"`
}

type GetMedOrder struct {
	ID uint `uri:"id" binding:"required"`
}

type UpdateMedOrder struct {
	Status        string `json:"status"`
	RansonScore   int    `json:"ranson_score"`
	MortalityRisk string `json:"mortality_risk"`
}

type FormMedOrder struct {
	ID uint `json:"id" binding:"required"`
	//CreatorID uint   `json:"creator_id" binding:"required"`
	//Password  string `json:"password" binding:"required"`
}

type EndOrCancelMedOrder struct {
	ID     uint   `uri:"id" binding:"required"`
	Status string `uri:"status" binding:"required,oneof=completed rejected"`
}

type GetModerator struct {
	ModeratorID uint   `form:"moderator_id" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type DeleteMedOrder struct {
	ID uint `json:"id" binding:"required"`
}
