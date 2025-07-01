package dto

type CreateAccountRequest struct {
	Login string `json:"login" binding:"required"`
}
