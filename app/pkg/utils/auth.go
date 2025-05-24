package utils

import "regexp"

func ValidatePhoneNumber(phoneNumber string) bool {
	starting := `^\+998\d{9}$`
	correct, err := regexp.MatchString(starting, phoneNumber)
	if err != nil {
		return false
	}
	if !correct {
		return false
	}
	return true
}
