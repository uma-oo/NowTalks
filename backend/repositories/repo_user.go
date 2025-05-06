package repositories

import (
	"fmt"

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


// function to check if a specific item is there based on a specific value
// generic somehow
// we need to specify  the type aftewards ;)
// it will be used for the nickname , session and also the email checking
func (appRep *AppRepository) GetItem(typ string, field string, value string) ([]any, bool, *models.ErrorJson) {
	data := make([]any, 0)
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
