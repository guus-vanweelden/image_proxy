package filter

import (
	"image"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/guus-vanweelden/image_proxy/cache"
)

func invertCalculation(i image.Image) image.Image {
	return imaging.Invert(i)
}

func InvertImage(w http.ResponseWriter, r *http.Request) {
	urlString := r.FormValue("url")
	// log.Printf("convert Image URL: %s", urlString)

	if img, err := cache.Get(calcCacheKey(urlString, "INVERT", "")); err == nil {
		log.Print("Response cached result")
		responseImage(w, img, getWebpSupport(r))
		return
	}
	log.Printf("Hash: %s", calcCacheKey(urlString, "INVERT", ""))

	img, err := getImage(urlString)
	if err != nil {
		// log.Printf("Error getImage(): %s", err)
		responseError(err, w)
		return
	}
	img = invertCalculation(img)
	cache.Set(calcCacheKey(urlString, "INVERT", ""), img)
	responseImage(w, invertCalculation(img), getWebpSupport(r))
}
