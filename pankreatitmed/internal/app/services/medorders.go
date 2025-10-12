package services

import (
	"fmt"
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/mapper"
	"strconv"
	"time"
	//"pankreatitmed/internal/app/dto/request"
	//"pankreatitmed/internal/app/singleton"
	//"time"
	"errors"
)

type MedOrdersService interface {
	GetDraft(creatorID uint) (*response.SendCartMedOrder, error)
	List(status string, start, end time.Time) ([]response.SendMedOrders, error)
	Get(ID uint) (response.SendMedOrder, error)
	Update(ID uint, in *request.UpdateMedOrder) error
	Form(ID uint) error
	CancelOrEnd(ID, moderator uint, password, status string) error
	Delete(ID uint) error
}

type medOrdersService struct {
	repo MedOrdersRepoPort
}

func NewMedOrdersService(repo MedOrdersRepoPort) MedOrdersService {
	return &medOrdersService{repo: repo}
}

// TODO перенести сюда singleton из хэндлера
func (s *medOrdersService) GetDraft(creatorID uint) (*response.SendCartMedOrder, error) {
	o, err := s.repo.GetOrCreateDraftMedOrder(creatorID)
	if err != nil {
		return nil, err
	}
	amnt, err := s.repo.CountItems(o.ID)
	if err != nil {
		return nil, err
	}
	res := mapper.MedOrderToSendMedOrder(o, uint(amnt))
	fmt.Println(res)
	return &res, err
}

func (s *medOrdersService) List(status string, start, end time.Time) ([]response.SendMedOrders, error) {
	morders, err := s.repo.GetMedOrders(status, start, end)
	if err != nil {
		return nil, err
	}
	res := mapper.MedOrdersToSendMedOrders(morders)
	return res, nil
}

func (s *medOrdersService) Get(ID uint) (response.SendMedOrder, error) {
	o, items, err := s.repo.GetMedOrderWithItems(ID)
	res := mapper.MedOrderToSendMedOrderWithItems(o, items)
	return res, err
}

func (s *medOrdersService) Update(ID uint, in *request.UpdateMedOrder) error {
	return s.repo.UpdateMedOrder(ID, in)
}

// TODO сделать проверку на соответстиве пользователя и создателя, а так же проверка на черновик
func (s *medOrdersService) Form(ID uint) error {
	check, err := s.repo.IsMedOrderDraft(ID)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("MedOrderIsNotDraft")
	}
	return s.repo.FormMedOrder(ID)
}

func (s *medOrdersService) CancelOrEnd(ID, moderator uint, password, status string) error {
	check, err := s.repo.IsMedOrderFormed(ID)
	if err != nil {
		return err
	}
	if !check {
		return errors.New("MedOrderIsNotFormed")
	}
	_, criteria, err := s.repo.GetMedOrderWithItems(ID)
	if err != nil {
		return err
	}
	if CheckReadyToCanselOrEnd(criteria) {
		rans, rsk, err := s.computeRanson(criteria)
		if err != nil {
			return err
		}
		if err := s.repo.SetRansonAndRisk(ID, rans, rsk); err != nil {
			return err
		}
		if err := s.repo.EndOrCancelMedOrder(ID, moderator, status); err != nil {
			return err
		}
	} else {
		return errors.New("Not all value fields are complete")
	}
	return nil
}

func CheckReadyToCanselOrEnd(items []ds.MedOrderItem) bool {
	for _, item := range items {
		if item.ValueNum == nil {
			return false
		}
	}
	return true
}

func (s *medOrdersService) computeRanson(items []ds.MedOrderItem) (int, string, error) {
	var score int
	for _, item := range items {
		crit, err := s.repo.GetCriterionByID(item.CriterionID)
		if err != nil {
			return 0, "", err
		}

		if crit.RefHigh != nil {
			if *item.ValueNum > *crit.RefHigh {
				score++
			}
		} else if crit.RefLow != nil {
			if *item.ValueNum < *crit.RefLow {
				score++
			}
		} else {
			return 0, "", errors.New("One of Ref is null")
		}
	}
	println(score)
	return score, strconv.Itoa(score*100/11) + "%", nil
}

func (s *medOrdersService) Delete(ID uint) error {
	return s.repo.SoftDeleteOrderSQL(ID)
}
