package request

type GetCriterion struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
}

type GetCriteria struct {
	Query string `json:"query"`
}

type CreateCriterion struct {
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Duration    string  `json:"duration" binding:"required"`
	HomeVisit   bool    `json:"home_visit"`
	Status      string  `json:"status" binding:"omitempty, oneof=active deleted"`
	Unit        string  `json:"unit"`
	RefLow      float64 `json:"ref_low"`
	RefHigh     float64 `json:"ref_high"`
}

type UpdateCriterion struct {
	Code        *string  `json:"code"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Duration    *string  `json:"duration"`
	HomeVisit   *bool    `json:"home_visit"`
	Unit        *string  `json:"unit"`
	RefLow      *float64 `json:"ref_low"`
	RefHigh     *float64 `json:"ref_high"`
}

type DeleteCriterion struct {
	ID uint `json:"id" binding:"required"`
}

type AddToDraft struct {
	ID uint `json:"id" binding:"required"`
}

type CreateCriterionIamage struct {
	ID uint `json:"id"`
}
