package handlers

import (
	"beginner/myapi/models"
	"beginner/myapi/utils"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: "Ramesh", Email: "ramesh@gmail.com"}
	utils.WriteJSON(w, http.StatusOK, user)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req models.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]models.ErrorResponseMessage{
			"error": {
				Success: false,
				Message: "Invalid request body",
			},
		})
		return
	}

	// validation

	// if req.Email != "admin" || req.Password != "1234" {
	// 	utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
	// 		"error": "invalid credentials",
	// 	})
	// 	return
	// }

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"Name":  req.Name,
		"Email": req.Email,
	})

}
