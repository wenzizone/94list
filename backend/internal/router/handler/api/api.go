package api

import (
	"encoding/json"
	"fmt"
	"github.com/wenzizone/94list/backend/internal/log"
	"net/http"
)

type getlist struct {
	action   string `json:"type"`
	shorturl string `json:"shorturl"`
	dir      string `json:"dir"`
	root     string `json:"root"`
	pwd      string `json:"pwd"`
	page     int    `json:"page"`
	num      int    `json:"num"`
	order    string `json:"order"`
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	uPath := r.URL.Path
	uHeader := r.Header

	r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	json.Unmarshal(r.Form, &getlist{})
	// 注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据
	log.Infof("form content: %s", r.Form) // 这些信息是输出到服务器端的打印信息

	log.Infof("path: %s\n, header: %s\n", uPath, uHeader)
	fmt.Fprintf(w, "path: %s\nheader: %s\nform content: %s\n", uPath, uHeader, r.Form)
}

func GetSign(w http.ResponseWriter, r *http.Request) {
	uPath := r.URL.Path
	uHeader := r.Header

	log.Infof("path: %s, header: %s", uPath, uHeader)
	fmt.Fprintf(w, "path: %s", uPath)
}
