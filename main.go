package main

import (
	"fmt"
	"net/http"

	"github.com/guus-vanweelden/image_proxy/router"
)

func main() {
	http.Handle("/", router.Router())
	http.ListenAndServe(fmt.Sprintf(":%d", 8123), nil)
}
