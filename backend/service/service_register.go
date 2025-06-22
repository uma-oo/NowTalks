package service

import (
	"fmt"
	"regexp"

	"real-time-forum/backend/models"
	"real-time-forum/backend/utils"
)

// CHECK THE info of the user
func (s *AppService) Register(user *models.User) *models.ErrorJson {
	var registerErr models.RegisterError
	// check for the nickname and email

	if err := s.IsValidNickname(user.Nickname); err != nil {
		registerErr.Nickname = err.Error()
	}

	if errEMail := s.EmailVerification(user.Email); errEMail != nil {
		registerErr.Email = errEMail.Error()
	}

	if user.Age < 15 {
		registerErr.Age = "you are too Young!! Go play outside :)"
	}
	if user.Age >= 100 {
		registerErr.Age = "you need to rest :( "
	}
	if err := utils.FirstLastNameVerf(user.FirstName); err != nil {
		registerErr.FirstName = err.Error()
	}
	if errLast := utils.FirstLastNameVerf(user.LastName); errLast != nil {
		registerErr.LastName = errLast.Error()
	}
	if err := utils.PwdFormatVerf(user.Password); err != nil {
		registerErr.Password = err.Error()
	}
	if !utils.PwdVerification(user.Password, user.VerifPassword) {
		registerErr.VerifPassword = "passwords does not match!"
	}
	if !utils.CheckGender(user.Gender) {
		registerErr.Gender = "please be sure to enter Male or Female"
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

func (s *AppService) IsValidNickname(nickname string) error {
	if len(nickname) < 3 {
		return fmt.Errorf("username is too short")
	}
	if len(nickname) > 30 {
		return fmt.Errorf("username is too long")
	}
	usernameRegex := `^[a-zA-Z0-9_.-]+$`
	if match, _ := regexp.MatchString(usernameRegex, nickname); !match {
		return fmt.Errorf("username can only contain letters, digits, underscores, dots, and hyphens")
	}
	_, has_nickname, _ := s.repo.GetItem("users", "nickname", nickname)
	if has_nickname {
		return fmt.Errorf("username already exists")
	}
	return nil
}

func (s *AppService) EmailVerification(email string) error {
	if len(email) > 255 {
		return fmt.Errorf("email too long")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,20}$`
	if match, _ := regexp.MatchString(emailRegex, email); !match {
		return fmt.Errorf("invalid email format")
	}
	_, has_email, _ := s.repo.GetItem("users", "email", email)
	if has_email {
		return fmt.Errorf("email already in use")
	}
	return nil
}
