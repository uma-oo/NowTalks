package repositories

import (
	models "real-time-forum/backend/models"
)

// DB wash create user hya register hnayaa wlla hadak service aykllf ???  ;(
// hadshi taaafh
// y9dr ay wa7d ydiiruuu

func (appRep *AppRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users VALUES (?,?,?,?,?,?,?)`
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
