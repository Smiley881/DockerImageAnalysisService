package hndlr

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"project8/services"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultByte)
	log.Printf("%s %s - 200 OK", r.Method, r.URL.Path)
}

func NotFoundUrl(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s - 404 Not Found: %s", r.Method, r.URL.Path, "Введенный URL не существует")
	http.Error(w, "Введенный URL не существует", http.StatusNotFound)
}
