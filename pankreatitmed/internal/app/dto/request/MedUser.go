package request

type MedUserRegistration struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetMedUser struct {
	Login string `json:"login" binding:"required"`
}

type UpdateMedUser struct {
	Login       string `json:"login" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewLogin    string `json:"login" binding:"required"`
	NewPassword string `json:"password" binding:"required"`
}

type AuthenticateUser struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
