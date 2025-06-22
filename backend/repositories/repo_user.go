package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	models "real-time-forum/backend/models"
)

// DB wash create user hya register hnayaa wlla hadak service aykllf ???  ;(
// hadshi taaafh
// y9dr ay wa7d ydiiruuu

func (appRep *AppRepository) CreateUser(user *models.User) *models.ErrorJson {
	query := `INSERT INTO users (nickname, age, gender, firstName, lastName, email, password) VALUES (?,?,?,?,?,?,?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	defer stmt.Close()
	if _, err = stmt.Exec(strings.ToLower(user.Nickname), user.Age, strings.ToLower((user.Gender)), user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// TO GET THE USERS

func (appRep *AppRepository) GetUsers(offset, user_id int) ([]models.User, *models.ErrorJson) {
	var users []models.User
	query := `with 
	cte_latest_interaction as (
	SELECT
        CASE 
            WHEN senderID = ? THEN receiverID 
            ELSE senderID 
        END AS userID,
        MAX(createdAt) AS lastInteraction,
        message
    FROM messages
    WHERE senderID = ? OR receiverID = ?
    GROUP BY userID),
	cte_ordered_users as (
	SELECT i.message, coalesce(i.lastInteraction, 0) as lastInteraction  , u.userID, u.nickname 
	from users u 
		Left JOIN cte_latest_interaction i 
		ON i.userID = u.userID
		WHERE u.userID != ?
	),
	cte_notifications as (
		select 
		senderID,
		count(*) as notifications 
	from
		messages 
	WHERE
		readStatus = 0
		And receiverID = ?
	GROUP BY
		senderID
	)
	SELECT 
		u.userID, u.nickname , coalesce(u.message, ""), u.lastInteraction, coalesce(n.notifications,0)
	FROM cte_ordered_users u
		LEFT JOIN cte_notifications n ON u.userID = n.senderID
		ORDER BY u.lastInteraction DESC,
		u.nickname
; `
	rows, err := appRep.db.Query(query, user_id, user_id, user_id, user_id, user_id)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Nickname, &user.LastMessage, &user.LastInteraction, &user.Notfications); err != nil {
			return nil, &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v ", err)}
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
				LoginField: "invalid login credentials!",
				Password:   "invalid login credentials!",
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

func (appRepo *AppRepository) UserExists(id int) (bool, *models.ErrorJson) {
	var exists bool
	query := ` SELECT EXISTS(SELECT 1 FROM users WHERE userID = ?);`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return false, models.NewErrorJson(500, fmt.Sprintf("%v 3", err))
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, &models.ErrorJson{Status: 400, Message: "user not found"}
	}
	return exists, nil
}
