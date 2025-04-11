package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := app.renderTemplate(w, nil)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
}

func (app *App) apiHandler(w http.ResponseWriter, r *http.Request) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		respondWithError(w, http.StatusUnsupportedMediaType, "Content Type is not application/json")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&app.key)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			respondWithError(w, http.StatusBadRequest, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field)
		} else {
			respondWithError(w, http.StatusBadRequest, "Bad Request "+err.Error())
		}
		return
	}
	err = app.key.Press()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(nil)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
}

func (app *App) staticFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, 405, "Method Not Allowed")
		return
	}
	http.ServeFile(w, r, app.staticPath+strings.TrimPrefix(r.URL.Path, "/static"))
}

func (app *App) renderTemplate(w http.ResponseWriter, data any) error {
	if app.debugMode {
		if err := app.initTemplates(); err != nil {
			return err
		}
	}
	return app.cachedTemplates.Execute(w, data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}
