package services

import "fmt"

type MedOrderItemsService interface {
	Delete(medorder, criterion uint) error
	Update(medorder, criterion uint, position *uint, val *float64) error
}

type medOrderItemsService struct {
	repo MedOrderItemsRepoPort
}

func NewMedOrderItemsService(repo MedOrderItemsRepoPort) MedOrderItemsService {
	return &medOrderItemsService{repo: repo}
}
func (s *medOrderItemsService) Delete(medorder, criterion uint) error {
	fmt.Println(medorder, criterion)
	fmt.Println(s.repo)
	return s.repo.DeleteFromOrder(medorder, criterion)
}

func (s *medOrderItemsService) Update(medorder, criterion uint, position *uint, val *float64) error {
	return s.repo.UpdateMedOrderItem(medorder, criterion, position, val)
}
