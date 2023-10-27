package api

import (
  "fmt"
  "github.com/wenzizone/94list/backend/internal/log"
  "net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
  uPath := r.URL.Path
  uHeader := r.Header

  log.Infof("path: %s, header: %s", uPath, uHeader)
  fmt.Fprintf(w, "path: %s", uPath)
}

func GetSign(w http.ResponseWriter, r *http.Request) {
  uPath := r.URL.Path
  uHeader := r.Header

  log.Infof("path: %s, header: %s", uPath, uHeader)
  fmt.Fprintf(w, "path: %s", uPath)
}
