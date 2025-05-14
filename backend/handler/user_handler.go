package handler

// func (h *AppHandler) addUser(w http.ResponseWriter, r http.Request) {
// 	var user models.User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		errJson := models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)}
// 		WriteJsonErrors(w, errJson)
// 		return
// 	}
// 	err := h.service.AddUser(&user)
// 	if err != nil {
// 		WriteJsonErrors(w, *err)
// 		return
// 	}
// }
