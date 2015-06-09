package filter

import (
	"image"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
)

func resizeCalculation(i image.Image, width int, height int) image.Image {
	return imaging.Resize(i, width, height, imaging.Box)
}

func ResizeImage(w http.ResponseWriter, r *http.Request) {
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
