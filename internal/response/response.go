package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

const (
	statusOk  = "ok"
	statusErr = "failed"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "aplication.json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&data)
}

func GeneralError(err string) Response {
	return Response{
		Status: statusErr,
		Error:  err,
	}
}

func Success(data interface{}) Response {
	return Response{
		Status: statusOk,
		Data:   data,
	}
}

func ValidationErr(errs validator.ValidationErrors) Response {

	var erMsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			erMsg = append(erMsg, fmt.Sprintf("field %s is required", err.Field()))
		default:
			erMsg = append(erMsg, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}

	return Response{
		Status: statusErr,
		Error:  strings.Join(erMsg, ", "),
	}

}
