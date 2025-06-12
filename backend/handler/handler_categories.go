package handler

import (
	"net/http"

	"real-time-forum/backend/models"
)




// Require authentication ??? 
func (catHandler *CategoriesHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		WriteJsonErrors(w, *models.NewErrorJson(405, "ERROR!! Method Not Allowed!"))
		return
	}
	categories, err := catHandler.service.GetAllCategories()
	if err != nil {
		WriteJsonErrors(w, *models.NewErrorJson(err.Status, err.Message))
		return
	}
	WriteDataBack(w, categories)
}



