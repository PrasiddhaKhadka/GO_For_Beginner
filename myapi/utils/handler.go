package utils

import "net/http"

type AppHandler func(w http.ResponseWriter, r *http.Request) error

// Converts AppHandler into a standard http.HandlerFunc
func Handle(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			// Check if it's our AppError
			if appErr, ok := err.(*AppError); ok {
				WriteJSON(w, appErr.Status, map[string]string{
					"error": appErr.Message,
				})
				return
			}
			// Unknown error — don't leak internals to client
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
