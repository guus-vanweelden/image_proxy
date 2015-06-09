package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guus-vanweelden/image_proxy/filter"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/gray", filter.GrayImage).Methods("GET")
	r.HandleFunc("/invert", filter.InvertImage).Methods("GET")
	r.HandleFunc("/blur", filter.BlurImage).Methods("GET")
	r.HandleFunc("/resize", filter.ResizeImage).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}
