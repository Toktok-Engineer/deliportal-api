package model

type LoginParameter struct {
	Username string `json:"username" binding:"required" validate:"username"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User         User
	AccessToken  string
	RefreshToken string
}
