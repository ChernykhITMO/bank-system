package dto

type TransferRequest struct {
	Id1    string  `json:"id1" binding:"required"`
	Id2    string  `json:"id2" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}
