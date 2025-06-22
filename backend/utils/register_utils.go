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
	if len(password) < 6 {
		return fmt.Errorf("password is too short")
	}
	if len(password) > 64 {
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
		fmt.Errorf("the name is too short")
	}
	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		fmt.Errorf("the name must only contain letters ")
	}

	return nil 
}
