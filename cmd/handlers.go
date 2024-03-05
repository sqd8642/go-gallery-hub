package main

import (
	"encoding/json"
	"github.com/sqd8642/go-gallery-hub/pkg/model" 
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *application) createImageHandler(w http.ResponseWriter, r *http.Request) { 
	var input struct {
		Url            string `json:"Url"`
		Caption        string `json:"Caption"`
		
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	image := &model.Image{ 
		Url:          input.Url,
		Caption:      input.Caption
	}

	err = app.models.Images.Insert(menu)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, menu)
}

func (app *application) getImageHandler(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r)
	param := vars["imageId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	image, err := app.models.Images.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, image)
}

func (app *application) updateMenuHandler(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r)
	param := vars["imageId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	menu, err := app.models.Images.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		Caption          *string `json:"caption"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if image.Title != nil {
		image.Caption = *input.Caption
	}
	err = app.models.Images.Update(image)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, menu)
}

func (app *application) deleteImageHandler(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r)
	param := vars["imageId"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid menu ID")
		return
	}

	err = app.models.Menus.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error { 
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}