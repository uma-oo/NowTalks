package service

import (
	"fmt"

	"real-time-forum/backend/models"
	"real-time-forum/backend/utils"
)

func (s *AppService) Register(user *models.User) *models.ErrorJson {
	var errJson models.ErrorJson
	var registerErr models.RegisterError
	// check for the nickname and email
	_, has_nickname, _ := s.repo.GetItem("users", "nickname", user.Nickname)
	_, has_email, _ := s.repo.GetItem("users", "email", user.Email)
	if has_nickname {
		registerErr.Nickname = "ERROR! Username already exists"
	}
	if has_email {
		registerErr.Email = "ERROR! Email already in use"
	}
	if !utils.IsValidNickname(user.Nickname) {
		registerErr.Nickname = "ERROR! Username Format is Incorrect"
	}
	if user.Age < 18 || user.Age >= 100 {
		registerErr.Age = "ERROR! Age had to be 18 and less than 100"
	}
	if !utils.FirstLastNameVerf(user.FirstName) {
		registerErr.FirstName = "ERROR! Sorry your First Name can't be stored on our system"
	}
	if !utils.FirstLastNameVerf(user.LastName) {
		registerErr.LastName = "ERROR! Sorry your Last Name can't be stored on our system"
	}
	if !utils.CheckEmailFormat(user.Email) {
		registerErr.Email = "ERROR! email Format is Incorrect"
	}
	if !utils.PwdFormatVerf(user.Password) {
		registerErr.Password = "ERROR! Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character"
	}
	if !utils.PwdVerification(user.Password, user.VerifPassword) {
		registerErr.VerifPassword = "ERROR! Passwords are not matched!"
	}
	// check if struct 3amra wlla la
	if registerErr != (models.RegisterError{}) {
		errJson.Status = 400
		errJson.Message = registerErr
		return &errJson
	}

	// hash the password here !! before the database insertion
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	} else {
		user.Password = hash
	}
    
	err = s.repo.CreateUser(user)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}

	return nil
}


