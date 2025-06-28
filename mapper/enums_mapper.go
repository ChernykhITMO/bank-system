package mapper

import (
	enums2 "bankSystem/domain/enums"
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

func EnumToString(sex enums2.Sex, hair enums2.Color) (string, string) {
	return string(sex), string(hair)
}
