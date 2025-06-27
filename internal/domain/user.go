package domain

import (
	"bankSystem/internal/domain/enums"
)

type User struct {
	Login     string
	Name      string
	Sex       enums.Sex
	HairColor enums.Color
	Friends   []string
	Accounts  []Account
}
