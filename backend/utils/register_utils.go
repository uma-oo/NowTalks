package utils

import (
	"fmt"
	"regexp"
)

// hadshi makhdmash
func CheckEmailFormat(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,20}$`
	if match, _ := regexp.MatchString(emailRegex, email); !match {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func PwdVerification(pwd string, pwdVerf string) bool {
	return pwd == pwdVerf
}

// Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character:

//	the regex "^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$"
//
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

// Sorry your name can't be stored on our system
func IsValidNickname(nickname string) bool {
	return regexp.MustCompile(`^[a-z0-9]+(?:[ _-][a-z0-9]+)*$`).MatchString(nickname)
}

// not sure
func FirstLastNameVerf(name string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+(?:[ -][A-Za-z]+)*$`).MatchString(name)
}
