package dto

type AuthRegisterPayload struct {
	Username string `json:"username" from:"username" binding:"required,min=3" example:""`
	Password string `json:"password" from:"password" binding:"required,min=3" example:""`
	Name     string `json:"name" from:"name" binding:"required,min=3" example:""`
}

type AuthRegisterResponse struct {
	Data    AuthRegisterDataResponse `json:"data"`
	Message string                   `json:"message" example:""`
}

type AuthRegisterDataResponse struct {
	Id        string `json:"id" example:""`
	Username  string `json:"username" example:""`
	Name      string `json:"name" example:""`
	CreatedAt string `json:"created_at" example:""`
	UpdatedAt string `json:"updated_at" example:""`
}
