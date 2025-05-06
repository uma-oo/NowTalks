package repositories

import (
	"fmt"
	"net/mail"
	"regexp"

	models "real-time-forum/backend/models"
)

// DB wash create user hya register hnayaa wlla hadak service aykllf ???  ;(
// hadshi taaafh
// y9dr ay wa7d ydiiruuu

func (appRep *AppRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (nickname, age, gender, firstName, lastName, email, password) VALUES (?,?,?,?,?,?,?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return nil
	}
	return nil
}

func (appRep *AppRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT userID,nickname, age, gender, firstName, lastName, email FROM users`
	rows, err := appRep.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Nickname, &user, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

// check if the nickname is duplicated

// function to check if a specific item is there based on a specific value

// generic somehow
// we need to specify  the type aftewards ;)
// it will be used for the nickname , session and also the email checking
func (appRep *AppRepository) GetItem(typ string, field string, value string) ([]interface{}, bool, *models.ErrorJson) {
	data := make([]interface{}, 0)
	query := fmt.Sprintf(`SELECT %v FROM %v WHERE %v=?`, field, typ, field)
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, false, models.NewErrorJson(500, "ERROR: Internal Server Error!!")
	}
	rows, err := stmt.Query(value)
	if err != nil {
		return nil, false, models.NewErrorJson(500, "ERROR: Internal Server Error!!")
	}
	for rows.Next() {
		var row any
		rows.Scan(&row)
		data = append(data, row)
	}
	if len(data) != 0 {
		return data, true, nil
	}
	return nil, false, nil
}

func CheckEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PwdVerification(pwd string , pwdVerf string) bool {
	return pwd==pwdVerf
}


// Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character:

//    the regex "^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$"

func PwdFormatVerf(password string) bool {
    reg , _ := regexp.Compile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	return reg.MatchString(password)
}

// Sorry your name can't be stored on our system

func FirstLastNameVerf() bool {
	return true
}
