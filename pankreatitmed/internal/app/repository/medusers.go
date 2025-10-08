package repository

import (
	"pankreatitmed/internal/app/ds"
	"pankreatitmed/internal/app/dto/request"
)

func (r *Repository) CreateMedUser(user *ds.MedUser) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetMedUserByLogin(login string) (*ds.MedUser, error) {
	var user ds.MedUser
	err := r.db.Where("login = ?", login).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) GetMedUserByID(id uint) (*ds.MedUser, error) {
	var user ds.MedUser
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) ChangeMedUser(id uint, user *request.UpdateMedUser) error {
	//println(*user.Login, *user.Password)
	return r.db.Model(&ds.MedUser{}).Where("id = ?", id).Updates(user).Error
}
