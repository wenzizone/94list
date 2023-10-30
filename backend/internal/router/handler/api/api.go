package api

import (
	"encoding/json"
	"fmt"
	"github.com/wenzizone/94list/backend/internal/log"
	"io"
	"net/http"
)

type request interface{}

type getlist struct {
	Action   string `json:"type"`
	Shorturl string `json:"shorturl"`
	Dir      string `json:"dir"`
	Root     string `json:"root"`
	Pwd      string `json:"pwd"`
	Page     string `json:"page"`
	Num      string `json:"num"`
	Order    string `json:"order"`
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var form getlist
	uPath := r.URL.Path
	uHeader := r.Header
	body, _ := io.ReadAll(r.Body)

	defer r.Body.Close()

	json.Unmarshal(body, &form)
	log.Infof("form content: %s", form) // 这些信息是输出到服务器端的打印信息

	log.Infof("path: %s\n, header: %s\n", uPath, uHeader)
	fmt.Fprintf(w, "path: %s\nheader: %s\nform content: %s\n", uPath, uHeader, r.Form)

}

func request(r request) map[string]string {

	return
}

func GetSign(w http.ResponseWriter, r *http.Request) {
	uPath := r.URL.Path
	uHeader := r.Header

	log.Infof("path: %s, header: %s", uPath, uHeader)
	fmt.Fprintf(w, "path: %s", uPath)
}
