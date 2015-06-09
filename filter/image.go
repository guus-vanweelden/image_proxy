package filter

import (
	"errors"
	"fmt"
	"image"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/disintegration/imaging"
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

func responseImage(w http.ResponseWriter, img image.Image, webp bool) {
	if webp {
		// log.Printf("Browser supports webP")
	}

	imaging.Encode(w, img, imaging.JPEG)
}

func responseError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("%s", err)))
}

func getWebpSupport(r *http.Request) bool {
	webp := false

	// log.Printf("Header: %+v", r.Header["Accept"])
	for _, accepts := range r.Header["Accept"] {
		if strings.Contains(accepts, "image/webp") {
			webp = true
			break
		}
	}
	return webp
}
