package dto

type DepositWithdrawRequest struct {
	Id     string  `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
