package app

import (
	"log/slog"
	"net/http"
)

func getHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("get request"))
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func postHandler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte("post request"))
		if err != nil {
			logger.Error(err.Error())
		}
	}
}
