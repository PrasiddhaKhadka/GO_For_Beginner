package handlers

import (
	"beginner/myapi/models"
	"beginner/myapi/utils"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: "Ramesh", Email: "ramesh@gmail.com"}
	utils.WriteJSON(w, http.StatusOK, user)
}
