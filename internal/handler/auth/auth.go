package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AnkitBishen/courseHub/internal/response"
	"github.com/AnkitBishen/courseHub/internal/storage"
	"github.com/AnkitBishen/courseHub/internal/stype"
	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Gconf = &oauth2.Config{
	ClientID:     "166365392915-p3c9f1a31katdvgied4lsi5o0kpemfo5.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

var oauthStateString = "randomstate" // Random string to protect against CSRF

func Register(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user stype.UserRegister

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
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
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
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
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
			return
		}

		// validate credentails
		res, err := storage.UserLoginValidation(user)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
			return
		}

		if res {
			response.WriteJson(w, http.StatusAccepted, response.Success([]string{"user validate successfully"}))
		} else {
			response.WriteJson(w, http.StatusAccepted, response.Success([]string{"Invalid credentails successfully"}))
		}

	}
}

func OauthLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		url := Gconf.AuthCodeURL(oauthStateString)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func OauthLoginCallback(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verify state
		state := r.FormValue("state")
		if state != oauthStateString {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError("State parameter does not match"))
			return
		}

		// Exchange authorization code for token
		code := r.FormValue("code")
		token, err := Gconf.Exchange(context.Background(), code)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError("Failed to exchange token: "+err.Error()))
			return
		}

		// Fetch user information
		client := Gconf.Client(context.Background(), token)
		res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError("Failed to fetch user info: "+err.Error()))
			return
		}
		defer res.Body.Close()

		if res.Status == "200 OK" {
			// add into Db
			var user stype.UserRegister
			err := json.NewDecoder(res.Body).Decode(&user)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
				return
			}

			res, err := storage.UserValidateFromOauth(user)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err.Error()))
				return
			}

			if res {
				response.WriteJson(w, http.StatusOK, response.Success("User validate successfully"))
				return
			}

		}

		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError("something went wrong"))

	}
}
