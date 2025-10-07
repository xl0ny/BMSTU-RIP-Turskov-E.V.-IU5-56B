// internal/app/service/criteria_service.go
package services

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/singleton"
)

// минимальный контракт репозитория, НУЖНЫЙ именно этому сервису

type CriteriaService interface {
	List(q string) ([]ds.Criterion, error)
	Get(id uint) (*ds.Criterion, error)
	Create(c *ds.Criterion) error
	Update(id uint, in *request.UpdateCriterion) error
	Delete(id uint) error
	ToDradt(id uint) error
}

type criteriaService struct {
	repo CriteriaRepoPort
}

func NewCriteriaService(repo CriteriaRepoPort) CriteriaService {
	return &criteriaService{repo: repo}
}

func (s *criteriaService) List(q string) ([]ds.Criterion, error) {
	return s.repo.GetCriteria(q)
}

func (s *criteriaService) Get(id uint) (*ds.Criterion, error) {
	return s.repo.GetCriterionByID(id)
}

func (s *criteriaService) Create(c *ds.Criterion) error {
	id, err := s.repo.GetSeq()
	if err != nil {
		return err
	}
	c.ID = id

	err = s.repo.CreateCriterion(c)
	if err != nil {
		s.repo.ResetCriterionSequence()
	}
	return err
}

func (s *criteriaService) Update(id uint, in *request.UpdateCriterion) error {
	return s.repo.UpdateCriterion(id, in)
}

func (s *criteriaService) Delete(id uint) error {
	return s.repo.DeleteCriterion(id)
}

func (s *criteriaService) ToDradt(id uint) error {
	oi, err := s.repo.GetOrCreateDraftMedOrder(singleton.GetCurrentUser().ID)
	println(oi)
	if err != nil {
		return err
	}
	return s.repo.AddItem(oi.ID, id)
}
