package service

import (
	"fmt"

	"real-time-forum/backend/models"
	"real-time-forum/backend/utils"
)

// CHECK THE info of the user
func (s *AppService) Register(user *models.User) *models.ErrorJson {
	var registerErr models.RegisterError
	// check for the nickname and email
	_, has_nickname, _ := s.repo.GetItem("users", "nickname", user.Nickname)
	_, has_email, _ := s.repo.GetItem("users", "email", user.Email)
	if has_nickname {
		registerErr.Nickname = "Username already exists"
	}
	if errEMail := utils.CheckEmailFormat(user.Email); errEMail != nil {
		registerErr.Email = errEMail.Error()
	}

	if has_email {
		registerErr.Email = "Email already in use"
	}
	if !utils.IsValidNickname(user.Nickname) {
		registerErr.Nickname = "Username Format is Incorrect"
	}
	if user.Age < 15 || user.Age >= 100 {
		registerErr.Age = "You are too Young!! Go play outside :)"
	}
	if !utils.FirstLastNameVerf(user.FirstName) {
		registerErr.FirstName = "Sorry your First Name can't be stored on our system"
	}
	if !utils.FirstLastNameVerf(user.LastName) {
		registerErr.LastName = "Sorry your Last Name can't be stored on our system"
	}
	if err := utils.PwdFormatVerf(user.Password); err != nil {
		registerErr.Password = fmt.Sprintf("ERROR!! %s", err.Error())
	}
	if !utils.PwdVerification(user.Password, user.VerifPassword) {
		registerErr.VerifPassword = "Passwords does not match!"
	}
	if !utils.CheckGender(user.Gender) {
		registerErr.Gender = "Please be sure to enter Male or Female"
	}
	// check if struct 3amra wlla la
	if registerErr != (models.RegisterError{}) {
		return &models.ErrorJson{Status: 400, Message: registerErr}
	}

	// hash the password here !! before the database insertion
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v just to verify 1", err)}
	} else {
		user.Password = hash
	}

	errJson := s.repo.CreateUser(user)
	if errJson != nil {
		return errJson
	}

	return nil
}
