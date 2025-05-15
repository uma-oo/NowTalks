package repositories

import (
	"database/sql"
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
	_, err = stmt.Exec(user.Id, session.Token, session.ExpDate)
	if err != nil {
		return &models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
	}
	return nil
}

// get the session by the user id or the user nickname !!

func (appRepo *AppRepository) GetSessionbyTokenEnsureAuth(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT userID, sessionToken , expiresAt FROM sessions WHERE sessionToken = ?`
	row := appRepo.db.QueryRow(query, token).Scan(&session.UserId, session.Token, session.ExpDate)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorizes Access"}
	}
	return &session, nil
}

func (appRepo *AppRepository) HasValidToken(token string) (bool, *models.Session) {
	session := models.NewSession()
	query := `SELECT userID, sessionToken , expiresAt FROM sessions WHERE sessionToken = ?`
	row := appRepo.db.QueryRow(query, token).Scan(&session.UserId, &session.Token, &session.ExpDate)

	if row == sql.ErrNoRows {
		return false, nil
	}

	if (session != &models.Session{}) {
		return true, session
	}
	return false, nil
}

func (appRep *AppRepository) GetUserSession(field any) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT * FROM sessions WHERE userID = ?`
	row := appRep.db.QueryRow(query, field)
	err := row.Scan(&session.Id, &session.UserId, &session.Token, &session.ExpDate)
	if err != nil {
		return nil, &models.ErrorJson{Status: 400, Message: "ERROR!! No session is set for this User"}
	}
	return &session, nil
}

func (appRep *AppRepository) UpdateSession(session *models.Session, new_session *models.Session) *models.ErrorJson {
	query := `UPDATE sessions SET sessionToken = ? , expiresAt = ? where sessionToken= ?`
	_, err := appRep.db.Exec(query, new_session.Token, new_session.ExpDate, session.Token)
	if err != nil {
		return models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	return nil
}

func (appRep *AppRepository) DeleteSession(session models.Session) *models.ErrorJson {
	query := `DELETE FROM sessions WHERE sessionToken = ?`
	_, err := appRep.db.Exec(query, session.Token)
	if err != nil {
		return models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	return nil
}
