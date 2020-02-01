package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ImageData struct {
	Data string `json:"data"`
}

func ResizeImage(w http.ResponseWriter, r *http.Request) {
	var imageData ImageData
	vars := mux.Vars(r)

	width := vars["width"]
	height := vars["height"]

	if err := json.NewDecoder(r.Body).Decode(&imageData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Desired width %s, desired height %s, data %s", width, height, imageData.Data)
}
