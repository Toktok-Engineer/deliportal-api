package model

type LoginParameter struct {
	Username string `json:"username" binding:"required" validate:"username"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
