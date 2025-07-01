package dto

type DeleteAccountRequest struct {
	Id    string `json:"id" binding:"required"`
	Login string `json:"login" binding:"required"`
}
