package repositories

import (
	"database/sql"
	"fmt"

	"real-time-forum/backend/models"
)

type AppRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{db: db}
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

	defer rows.Close()
	
	if len(data) != 0 {
		return data, true, nil
	}
	return nil, false, nil
}

func (appRepo *AppRepository) GetUserbyToken(token string) (*models.Session, *models.ErrorJson) {
	session := models.Session{}
	query := `SELECT userID, sessionToken , expiresAt FROM sessions WHERE sessionToken = ?`
	row := appRepo.db.QueryRow(query, token).Scan(&session.UserId, session.Token, session.ExpDate)
	if row == sql.ErrNoRows {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorizes Access"}
	}
	return &session, nil
}
