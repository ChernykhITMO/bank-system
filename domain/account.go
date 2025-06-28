package domain

type Account struct {
	Id      string
	Balance float64
	Login   string
	History []Transaction
}
