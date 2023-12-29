package pkg

import "errors"

type error_string struct {
	message string
}

func (error_str error_string) Error() string {
	return error_str.message
}

func NewErrorTest(text string) error {
	return &error_string{text}
}
func AgeTestError(age int) (int, error) {
	if age < 18 {
		return age, errors.New("Usia anda tidak cukup tua")
	} else {
		return age, nil
	}
}
