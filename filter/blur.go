package filter

import (
	"image"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
)

func blurCalculation(i image.Image, blur float64) image.Image {
	return imaging.Blur(i, blur)
}

func BlurImage(w http.ResponseWriter, r *http.Request) {
	urlString := r.FormValue("url")
	blurString := r.FormValue("blur")

	blurInt, err := strconv.Atoi(blurString)
	// log.Printf("Width: %s %d", widthString, width)
	if err != nil {
		blurInt = 100
	}
	blur := float64(blurInt) / 100

	img, err := getImage(urlString)
	if err != nil {
		//	log.Printf("Error getImage(): %s", err)
		responseError(err, w)
		return
	}

	responseImage(w, blurCalculation(img, blur), getWebpSupport(r))
}
