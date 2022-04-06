package model

type RenewToken struct {
	ID       uint64 `json:"id" binding:"required"`
	Username string `json:"username" binding:"required" validate:"username"`
}

type RenewTokenResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
