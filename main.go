package main

import (
	"ResizeImageTool/handlers"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func transformImage(path string, height, width uint) {
	reader, err := os.Open(path)
	CheckError(err)
	defer reader.Close()
	img, _, err := image.Decode(reader)
	CheckError(err)

	newImg := resize.Resize(width, height, img, resize.NearestNeighbor)
	out, err := os.Create(path + "random_name.jpg")
	CheckError(err)

	err = jpeg.Encode(out, newImg, &jpeg.Options{Quality: 100})
	CheckError(err)
}

func main() {
	//transformImage("andrea.jpeg", 1080, 1920)

	router := mux.NewRouter()
	router.HandleFunc("/resize/{width}/{height}", handlers.ResizeImage).Methods("POST")
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8181",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server is running at address %s: \n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
