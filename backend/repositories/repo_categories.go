package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

// Insert the categories (received as a slice of strings)
func (appRep *AppRepository) AddPostCategories(categoryIds []int, postId int) *models.ErrorJson {
	for _, id := range categoryIds {
		query := `INSERT INTO postCategories (categoryID, category) VALUES (? , ?)`
		stmt, err := appRep.db.Prepare(query)
		if err != nil {
			return models.NewErrorJson(500, fmt.Sprintf("%v", err))
		}
		defer stmt.Close()
		_, err = stmt.Exec(id, postId)
		if err != nil {
			return models.NewErrorJson(500, fmt.Sprintf("%v", err))
		}
	}
	return nil
}

func (appRepo *AppRepository) GetAllCategories() ([]models.Category, *models.ErrorJson) {
	categories := []models.Category{}
	query := `SELECT categoryID, category FROM categories`
	stmt, err := appRepo.db.Prepare(query)
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
	}
	defer rows.Close()
	for rows.Next() {
		category := models.Category{}
		if err := rows.Scan(&category.CategoryId, &category.CategoryName); err != nil {
			return nil, models.NewErrorJson(500, fmt.Sprintf("%v", err))
		}
		categories = append(categories, category)
	}

	return categories, nil
}



// Given 
func  (appRepo *AppRepository) GetPostCategories(post *models.Post) (*models.Post, *models.ErrorJson) {

 
}
