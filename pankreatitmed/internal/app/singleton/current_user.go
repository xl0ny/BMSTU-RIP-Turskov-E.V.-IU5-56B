package singleton

import "pankreatitmed/internal/app/ds"

var currentUser *ds.MedUser

func GetCurrentUser() *ds.MedUser {
	if currentUser == nil {
		currentUser = &ds.MedUser{
			ID:          1,
			Login:       "demo",
			Password:    "demo",
			IsModerator: false,
		}
	}
	return currentUser
}
