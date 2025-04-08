package hndlr

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	services "project8/pkg"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func BaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		if statusCode == http.StatusNotFound && r.URL.Path != "/api/v1/image-download-size" {
			log.Printf("%s %s - 404 Not Found: Страница не существует", r.Method, r.URL.Path)
		} else if statusCode == http.StatusOK {
			log.Printf("%s %s - 200 OK", r.Method, r.URL.Path)
		}
	})
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Printf("%s %s - 405 Method Not Allowed: %s", r.Method, r.URL.Path, "Недопустимый метод запроса")
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Printf("%s %s - 400 Bad Request: %s", r.Method, r.URL.Path, "Недопустимый запрос: Ожидается формат JSON")
		http.Error(w, "Ожидается формат JSON", http.StatusBadRequest)
		return
	}

	resultStruct, err := services.ImageDownloadSize(r.Body)
	if errors.Is(err, services.ErrInvalidFormatJson) || errors.Is(err, services.ErrInvalidfFormatInput) {
		log.Printf("%s %s - 400 Bad Request: %s", r.Method, r.URL.Path, "Недопустимый запрос: "+err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if errors.Is(err, services.ErrNotFound) {
		log.Printf("%s %s - 404 Not Found: %s", r.Method, r.URL.Path, err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resultByte, err := json.Marshal(&resultStruct)
	if err != nil {
		log.Printf("%s %s - 500 Internal Server Error: %s", r.Method, r.URL.Path, err.Error())
		http.Error(w, "Ошибка при выполнении запроса", http.StatusInternalServerError)
		return
	}

	w.Write(resultByte)

}
