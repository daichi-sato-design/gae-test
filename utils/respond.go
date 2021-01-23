package utils

import (
	"encoding/json"
	"net/http"

	"github.com/daichi-sato-design/gae-test/domain"
)

// Respond リクエスト成功時に構造体を返す関数です。
func Respond(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// RespondError リクエスト失敗時にApiErrを返す関数です。
func RespondError(w http.ResponseWriter, apiErr *domain.ApplicationError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.StatusCode)
	json.NewEncoder(w).Encode(apiErr)
}
