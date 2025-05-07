package service

import (
	"real-time-forum/backend/models"
	"strings"
)







func (s *AppService) Login(login *models.Login) *models.ErrorJson {
	var LoginERR = models.Login{}
	if strings.TrimSpace(login.LoginField)==""  {
      LoginERR.LoginField="ERROR! Login field is Empty!"
	}
	if strings.TrimSpace(login.Password)=="" {
		 LoginERR.Password="ERROR! Password field is Empty!"
	}

	// a finir tomorrow

	return nil
}
