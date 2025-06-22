package utils

import (
	"fmt"
	"regexp"
)

// hadshi makhdmash

func PwdVerification(pwd string, pwdVerf string) bool {
	return pwd == pwdVerf
}

// lookarounds are not possible
func PwdFormatVerf(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}
	if len(password) > 50 {
		return fmt.Errorf("password is too long")
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)
	if !hasLower || !hasUpper || !hasDigit || !hasSpecial {
		return fmt.Errorf("password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character")
	}
	return nil
}

func FirstLastNameVerf(name string) error {
	if len(name) < 3 {
		return fmt.Errorf("the name is too short")
	}
	if len(name) > 50 {
		return fmt.Errorf("the name is too long")
	}
	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		return fmt.Errorf("the name must only contain letters ")
	}

	return nil 
}
