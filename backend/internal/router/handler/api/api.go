package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wenzizone/94list/backend/internal/log"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const staticUrl = "https://pan.baidu.com"
const fileInfoBasedUrl = "https://pan.baidu.com/share/wxlist?channel=weixin&version=2.2.2&clienttype=25&web=1"

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
	root := 1
	//uPath := r.URL.Path
	//uHeader := r.Header

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		// 处理解析错误
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if form.Dir != "" {
		root = 0
	}
	//startUrl := fmt.Sprintf("%s/s/%s", staticUrl, form.Shorturl)
	//get(startUrl, nil)

	payloadValues := url.Values{}
	payloadValues.Set("shorturl", form.Shorturl)
	payloadValues.Set("dir", form.Dir)
	payloadValues.Set("root", strconv.Itoa(root))
	payloadValues.Set("pwd", form.Pwd)
	payloadValues.Set("page", "1")
	payloadValues.Set("num", "1000")
	payloadValues.Set("order", "time")
	payloadDataStr := payloadValues.Encode()
	payloadDataBytes := []byte(payloadDataStr)
	payloadBytesReader := bytes.NewReader(payloadDataBytes)

	headers := map[string]string{
		"Content-Type": "application/application/x-www-form-urlencoded",
	}

	//url := "https://pan.baidu.com/share/wxlist?channel=weixin&version=2.2.2&clienttype=25&web=1"
	post(fileInfoBasedUrl, payloadBytesReader, headers)

	//log.Infof("form content: %s", form) // 这些信息是输出到服务器端的打印信息

	//log.Infof("path: %s\n, header: %s\n", uPath, uHeader)
	//fmt.Fprintf(w, "path: %s\nheader: %s\nform content: %s\n", uPath, uHeader, r.Form)

}

func post(url string, payload io.Reader, headers map[string]string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		// handle error
		log.Infof("生成请求失败， %s", err)
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
	//return
}

func get(url string, payload io.Reader) {
	headers := map[string]string{
		"Referer":    "https://pan.baidu.com",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, payload)
	if err != nil {
		// handle error
		log.Infof("生成请求失败， %s", err)
		return
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}

func GetSign(w http.ResponseWriter, r *http.Request) {
	uPath := r.URL.Path
	uHeader := r.Header

	log.Infof("path: %s, header: %s", uPath, uHeader)
	fmt.Fprintf(w, "path: %s", uPath)
}
