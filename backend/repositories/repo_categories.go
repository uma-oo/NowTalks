package repositories

import (
	"fmt"

	"real-time-forum/backend/models"
)

// Insert the categories (received as a slice of INT)
func (appRep *AppRepository) AddPostCategories(post *models.Post, categories []any) (*models.Post, *models.ErrorJson) {
	for _, id := range categories {
		query := `INSERT INTO postCategories (categoryID, postID) VALUES (? , ?)
		RETURNING (SELECT category FROM categories 
		WHERE categories.categoryID = postCategories.categoryID);`
		stmt, err := appRep.db.Prepare(query)
		if err != nil {
			return nil, models.NewErrorJson(500, fmt.Sprintf("%v 1", err))
		}
		defer stmt.Close()
		var category string
		err = stmt.QueryRow(id, post.Id).Scan(&category)
		if err != nil {
			return nil, models.NewErrorJson(500, fmt.Sprintf("%v 2", err))
		}
		post.PostCategories = append(post.PostCategories, category)
	}
	return post, nil
}



// THIS is for the getCategories endpoint
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

