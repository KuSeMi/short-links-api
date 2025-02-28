package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, resp any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(resp)
}
