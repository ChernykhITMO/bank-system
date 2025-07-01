package mapper

import (
	enums2 "bankSystem/internal/domain/constants"
	"fmt"
	"strings"
)

func StringToEnum(sex, hair string) (enums2.Sex, enums2.Color, error) {
	sexLower := strings.ToLower(sex)
	hairLower := strings.ToLower(hair)

	var sexEnum enums2.Sex
	var hairEnum enums2.Color

	switch sexLower {
	case "male":
		sexEnum = enums2.SexMale
	case "female":
		sexEnum = enums2.SexFemale
	default:
		return "", "", fmt.Errorf("invalid sex: %s", sex)
	}

	switch hairLower {
	case "black":
		hairEnum = enums2.ColorBlack
	case "white":
		hairEnum = enums2.ColorWhite
	default:
		return "", "", fmt.Errorf("invalid hair color: %s", hair)
	}

	return sexEnum, hairEnum, nil
}
