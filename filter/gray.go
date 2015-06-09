package filter

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

func grayCalculation(i image.Image) image.Image {
	return imaging.Grayscale(i)
}

func GrayImage(w http.ResponseWriter, r *http.Request) {
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
