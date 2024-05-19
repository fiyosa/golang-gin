package dto

type AuthLoginPayload struct {
	Username string `json:"username" from:"username" binding:"required,min=3" example:""`
	Password string `json:"password" from:"password" binding:"required,min=3" example:""`
}

type AuthLoginResponse struct {
	Token string `json:"token" example:""`
}
