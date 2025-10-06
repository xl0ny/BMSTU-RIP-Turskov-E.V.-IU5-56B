package handler

import (
	"context"
	"pankreatitmed/internal/app/dto"
	"pankreatitmed/internal/app/dto/response"
)

type CriteriaService interface {
	List(ctx context.Context, q string) dto.List[response.SendCriterion]
	//Get(ctx context.Context, id uint) (*response.Criterion, error)
	//Create(ctx context.Context, in request.CreateCriterion) (uint, error)
	//Update(ctx context.Context, id uint, in request.UpdateCriterion) error
	//Delete(ctx context.Context, id uint) error
	//
	//AddToDraft(ctx context.Context, userID, criterionID uint) error
}
