package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "aplication.json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(&data)
}
