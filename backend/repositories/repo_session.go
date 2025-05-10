package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

func (appRep *AppRepository) CreateUserSession(session *models.Session, user *models.User) *models.ErrorJson {
	query := `INSERT INTO sessions (userID, sessionToken, expiresAt) VALUES (?,?,?)`
	stmt, err := appRep.db.Prepare(query)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	
	defer stmt.Close()
	fmt.Println("session", session.Token, session.ExpDate, user.Id)
	_, err = stmt.Exec(user.Id, session.Token, session.ExpDate)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// get the session by the user id or the user nickname !!

func (appRep *AppRepository) GetUserSession(field any) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT * FROM sessions WHERE userID = ?`
	row := appRep.db.QueryRow(query, field)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpDate)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Message: "error u safi"}
	}
	return &session, nil
}
