package services

import (
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/mapper"
	"pankreatitmed/internal/app/singleton"
)

type MedUsersService interface {
	Registrate(usr request.MedUserRegistration) error
	GetMyField() (*response.SendMedUserField, error)
	UpdateMyField(user *request.UpdateMedUser) error
}

type medUsersService struct {
	repo MedUsersRepoPort
}

func NewMedUsersService(repo MedUsersRepoPort) MedUsersService {
	return &medUsersService{repo: repo}
}

func (s *medUsersService) Registrate(usr request.MedUserRegistration) error {
	user := mapper.MedUserRegistrationToMedUser(usr)
	return s.repo.CreateMedUser(&user)
}

func (s *medUsersService) GetMyField() (*response.SendMedUserField, error) {
	user, err := s.repo.GetMedUserByID(singleton.GetCurrentUser().ID)
	if err != nil {
		return nil, err
	}
	res := mapper.MedUserToSendMedUserFields(user)
	return &res, nil
}

func (s *medUsersService) UpdateMyField(user *request.UpdateMedUser) error {
	id := singleton.GetCurrentUser().ID
	return s.repo.ChangeMedUser(id, user)
}
