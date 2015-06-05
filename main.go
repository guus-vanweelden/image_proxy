package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
)

func getImage(urlString string) (image.Image, error) {
	imageUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(imageUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("getImage(): Response != 200")
	}
	img, format, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Format: %s", format)
	return img, nil
}

func invertCalculation(i image.Image) image.Image {
	return imaging.Invert(i)
}

func responseImage(w http.ResponseWriter, img image.Image, webp bool) {
	if webp {
		// log.Printf("Browser supports webP")
	}

	imaging.Encode(w, img, imaging.JPEG)
}

func getWebpSupport(r *http.Request) bool {
	webp := false

	log.Printf("Header: %+v", r.Header["Accept"])
	for _, accepts := range r.Header["Accept"] {
		if strings.Contains(accepts, "image/webp") {
			webp = true
			break
		}
	}
	return webp
}

func responseError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("%s", err)))
}

func grayImage(w http.ResponseWriter, r *http.Request) {
	urlString := r.FormValue("url")
	// log.Printf("convert Image URL: %s", urlString)

	img, err := getImage(urlString)
	if err != nil {
		// log.Printf("Error getImage(): %s", err)
		responseError(err, w)
		return
	}

	responseImage(w, grayCalculation(img), getWebpSupport(r))
}

func convertImage(w http.ResponseWriter, r *http.Request) {
	urlString := r.FormValue("url")
	// log.Printf("convert Image URL: %s", urlString)

	img, err := getImage(urlString)
	if err != nil {
		// log.Printf("Error getImage(): %s", err)
		responseError(err, w)
		return
	}

	responseImage(w, invertCalculation(img), getWebpSupport(r))
}

func grayCalculation(i image.Image) image.Image {
	return imaging.Grayscale(i)
}

func resizeCalculation(i image.Image, width int, height int) image.Image {
	return imaging.Resize(i, width, height, imaging.Box)
}

func resizeImage(w http.ResponseWriter, r *http.Request) {
	urlString := r.FormValue("url")
	widthString := r.FormValue("width")
	heightString := r.FormValue("height")

	width, err := strconv.Atoi(widthString)
	// log.Printf("Width: %s %d", widthString, width)
	if err != nil {
		width = 640
	}

	height, err := strconv.Atoi(heightString)
	// log.Printf("Height: %s %d", heightString, height)
	if err != nil {
		height = 480
	}

	img, err := getImage(urlString)
	if err != nil {
		//	log.Printf("Error getImage(): %s", err)
		responseError(err, w)
		return
	}

	responseImage(w, resizeCalculation(img, width, height), getWebpSupport(r))
}

func router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/gray", grayImage).Methods("GET")
	r.HandleFunc("/invert", convertImage).Methods("GET")
	r.HandleFunc("/resize", resizeImage).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}

func main() {
	http.Handle("/", router())
	http.ListenAndServe(fmt.Sprintf(":%d", 8123), nil)
}
