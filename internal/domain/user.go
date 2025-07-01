package domain

import (
	"bankSystem/internal/domain/constants"
)

type User struct {
	Login     string
	Name      string
	Sex       constants.Sex
	HairColor constants.Color
	Friends   []string
	Accounts  []Account
}
