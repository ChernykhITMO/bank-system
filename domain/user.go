package domain

import (
	enums2 "bankSystem/domain/enums"
)

type User struct {
	Login     string
	Name      string
	Sex       enums2.Sex
	HairColor enums2.Color
	Friends   []string
	Accounts  []Account
}
