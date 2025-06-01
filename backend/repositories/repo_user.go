package repositories

import (
	"database/sql"
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
	defer stmt.Close()
	if _, err = stmt.Exec(user.Nickname, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return nil
	}
	return nil
}

// TO GET THE USERS

func (appRep *AppRepository) GetUsers(offset int) ([]models.User, *models.ErrorJson) {
	var users []models.User
	query := `SELECT userID , nickname FROM usersLIMIT 10 OFFSET ?`
	rows, err := appRep.db.Query(query, offset)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Nickname); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		}
		users = append(users, user)
	}
	defer rows.Close()

	return users, nil
}

// chosen_field ( it may be the nickname or the email )
// the query must not include the password entered by the user
func (appRep *AppRepository) GetUser(login *models.Login) (*models.User, *models.ErrorJson) {
	user := models.NewUser()
	query := `SELECT userID, nickname, password 
	FROM users where nickname=? OR email =? `
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	row := stmt.QueryRow(login.LoginField, login.LoginField)
	err = row.Scan(&user.Id, &user.Nickname, &user.Password)
	if err == sql.ErrNoRows {
		return nil, &models.ErrorJson{
			Status: 401,
			Message: models.Login{
				LoginField: "ERROR!! Username or Email does not exist! Or Password Incorrect!",
				Password:   "ERROR!! Username or Email does not exist! Or Password Incorrect!",
			},
		}
	}
	return user, nil
}

// get the username from the userId
func (appRep *AppRepository) GetUserNameById(user_id int) (string, *models.ErrorJson) {
	var username string
	query := `SELECT nickname FROM users WHERE userID = ?`
	err := appRep.db.QueryRow(query, user_id).Scan(&username)
	if err != nil {
		return "", models.NewErrorJson(500, fmt.Sprintf("%v 3", err))
	}
	return username, nil
}
