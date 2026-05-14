package utils

import "net/http"

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func Handler(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {

			if appErr, ok := err.(*AppError); ok {
				WriteJSON(w, appErr.Code, map[string]string{
					"error": appErr.Message,
				})
				return
			}

			WriteJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "something went wrong",
			})
		}
	}
}

//  SAME AS
// func HandleNormal (falias func(w http.ResponseWriter, r *http.Request) error){
// to use
// falias()
// }
