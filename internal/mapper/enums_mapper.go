package mapper

import (
	"bankSystem/internal/domain/constants"
	"fmt"
	"strings"
)

func StringToEnum(sex, hair string) (constants.Sex, constants.Color, error) {
	sexLower := strings.ToLower(sex)
	hairLower := strings.ToLower(hair)

	var sexEnum constants.Sex
	var hairEnum constants.Color

	switch sexLower {
	case "male":
		sexEnum = constants.SexMale
	case "female":
		sexEnum = constants.SexFemale
	default:
		return "", "", fmt.Errorf("invalid sex: %s", sex)
	}

	switch hairLower {
	case "black":
		hairEnum = constants.ColorBlack
	case "white":
		hairEnum = constants.ColorWhite
	default:
		return "", "", fmt.Errorf("invalid hair color: %s", hair)
	}

	return sexEnum, hairEnum, nil
}
