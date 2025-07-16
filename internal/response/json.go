package response

import (
	"encoding/json"
	"net/http"
	"s3-demo/s3-demo-go/internal/entity"
)

func JSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "internal response error", http.StatusInternalServerError)
	}
}

func ResponseJson(w http.ResponseWriter, code int, massage string) error {
	msg := *&entity.Response{
		Massage: massage,
	}
	JSON(w, code, msg)
	return nil
}
