package writejson

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any, code int) error {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}

	return nil
}