package auth

import (
	"encoding/json"
	"net/http"

	"github.com/AnkitBishen/courseHub/internal/response"
	"github.com/AnkitBishen/courseHub/internal/storage"
	"github.com/AnkitBishen/courseHub/internal/stype"
	"github.com/go-playground/validator/v10"
)

func Register(storage storage.Storage) http.HandlerFunc {
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
		_, err = storage.UserRegister(user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// send success when registration is done
		response.WriteJson(w, http.StatusAccepted, response.Success([]string{"User registered successfully"}))
	}
}

func BasicAuth(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user stype.UserRegister

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validate credentails
		res, err := storage.UserLoginValidation(user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if res {
			response.WriteJson(w, http.StatusAccepted, response.Success([]string{"user validate successfully"}))
		} else {
			response.WriteJson(w, http.StatusAccepted, response.Success([]string{"Invalid credentails successfully"}))
		}

	}
}
