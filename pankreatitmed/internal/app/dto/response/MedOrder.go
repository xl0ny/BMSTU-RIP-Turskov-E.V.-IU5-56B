package response

import "time"

type SendCartMedOrder struct {
	MedORderId     string `json:"med_order_id"`
	CriteriaAmount string `json:"criteria_amount"`
}

type SendMedOrders struct {
	ID            uint      `json:"id"`
	Status        string    `json:"status"`
	FormedAt      time.Time `json:"formed_at"`
	FinishedAt    time.Time `json:"finished_at"`
	RansonScore   int       `json:"ranson_score"`
	MortalityRisk string    `json:"mortality_risk"`
}

type SendMedOrder struct {
	ID            uint            `json:"id"`
	Status        string          `json:"status"`
	CreatorID     uint            `json:"creator_id"`
	FormedAt      time.Time       `json:"formed_at"`
	FinishedAt    time.Time       `json:"finished_at"`
	ModeratorID   uint            `json:"moderator_id"`
	RansonScore   int             `json:"ranson_score"`
	MortalityRisk string          `json:"mortality_risk"`
	Criteria      []SendCriterion `json:"criteria"`
}
