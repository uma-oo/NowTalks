package service

import (
	"fmt"
	"strings"

	"real-time-forum/backend/models"
	"real-time-forum/backend/utils"
)

// so we need to check that the password entered from the API is the same stored in the database !!

func (s *AppService) Login(login *models.Login) *models.ErrorJson {
	LoginERR := models.Login{}

	if strings.TrimSpace(login.LoginField) == "" {
		LoginERR.LoginField = "ERROR! Login field is Empty!"
	}
	if strings.TrimSpace(login.Password) == "" {
		LoginERR.Password = "ERROR! Password field is Empty!"
	}
	//
	if LoginERR != (models.Login{}) {
		return &models.ErrorJson{Status: 400, Message: LoginERR}
	}

	
	// we need to check also if the user has the 401 error
	// check if the password and the login are wrong both
	user, err := s.repo.GetUser(login)
	if err != nil {
		switch err.Status {
		case 401:
			return err

		default:
			return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
	}
	// if only the password
	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return &models.ErrorJson{
			Status: 401,
			Message: models.Login{
				LoginField: "ERROR!! Username or Email does not exist! Or Password Incorrect!",
				Password:   "ERROR!! Username or Email does not exist! Or Password Incorrect!",
			},
		}
	}
	return nil
}
