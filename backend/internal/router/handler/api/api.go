package api

import (
	"fmt"
	"net/http"
)

type getList struct{}

func (gl *getList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uPath := r.URL.Path
	uHeader := r.Header

	fmt.Fprintf(w, "path: %s, heder: %s", uPath, uHeader)
}
