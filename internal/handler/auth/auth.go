package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AnkitBishen/courseHub/internal/response"
	"github.com/AnkitBishen/courseHub/internal/stype"
	"github.com/go-playground/validator/v10"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user stype.UserRegister

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validate inputs
		err = validator.New().Struct(user)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErr(validationErrors))
			return
		}

		// register the user
		// fmt.Println(os.Getenv("DB_Name"))

		// send success when registration is done
		response.WriteJson(w, http.StatusAccepted, response.Success(nil))
	}
}

// func Login() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		json.NewDecoder(r.Body).Decode(&)
// 		if err != nil {
// 			response.WriteJson(w, http.StatusBadRequest, nil)
// 			return
// 		}
// 		// send success when authentication is done
// 		response.WriteJson(w, http.StatusAccepted, nil)
// 	}
// }
