package utils

import (
	"net/mail"
	"regexp"
)

// hadshi makhdmash
func CheckEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PwdVerification(pwd string, pwdVerf string) bool {
	return pwd == pwdVerf
}

// Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character:

//	the regex "^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$"
//
// lookarounds are not possible
func PwdFormatVerf(password string) bool {
	if len(password) < 8 {
		return false
	}
	reg, err := regexp.Compile(`^[A-Za-z\d@$!%*?&]{8,}$`)
	if err != nil || !reg.MatchString(password) {
		return false
	}
	if !regexp.MustCompile("[a-z]{1,}").MatchString(password) {
		return false
	}
	if !regexp.MustCompile("[A-Z]{1,}").MatchString(password) {
		return false
	}
	if !regexp.MustCompile("[!@#$%^&*]{1,}").MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`\d{1,}`).MatchString(password) {
		return false
	}
	return true
}

// Sorry your name can't be stored on our system
func IsValidNickname(nickname string) bool {
	return regexp.MustCompile(`^[a-z0-9]+(?:[ _-][a-z0-9]+)*$`).MatchString(nickname)
}

// not sure
func FirstLastNameVerf(name string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+(?:[ -][A-Za-z]+)*$`).MatchString(name)
}
