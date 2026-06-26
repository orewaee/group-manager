package http

import (
	"encoding/json"
	"io"
	"net/http"
)

const maxBodySize = 1 << 20 // 1 MB

func read[T any](r *http.Request) (T, error) {
	var zero T
	r.Body = http.MaxBytesReader(nil, r.Body, maxBodySize)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return zero, err
	}

	var entity T
	if err := json.Unmarshal(data, &entity); err != nil {
		return zero, err
	}

	return entity, nil
}

func writeJson(w http.ResponseWriter, code int, dto any) {
	data, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

func writeError(w http.ResponseWriter, code int, message string) {
	writeJson(w, code, &ErrorResponse{
		Message: message,
	})
}
