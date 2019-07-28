package controllers

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/GoYoutubeServer/server/config"
	"github.com/GoYoutubeServer/server/models"
	"github.com/GoYoutubeServer/server/youtube"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Print(fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)))
}

func YoutubePage(w http.ResponseWriter, r *http.Request) {
	apiClient := youtube.New(config.Config.APIKey)
	v, err := apiClient.GetYoutube(r)
	if err != nil {
		fmt.Println("handler.go :", err)
		return
	}
	prettyJson, _ := json.MarshalIndent(v, "", " ")
	setHeadersHandler(w, r)
	log.Print(w.Write(prettyJson))
}
func Mock(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Path, r.URL.Query())

	m, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println("YoutubePageJson", m)

	fmt.Println("リクエスト:", r.RemoteAddr, r.Header.Get("Content-Type"))

	//ファイルを読み込み
	raw, err := ioutil.ReadFile("./server/models/mockYoutubeListSearch.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//モデルと対応
	var yj models.YoutubeJson

	//Unmarshal(JSON → Go Object)
	if err := json.Unmarshal(raw, &yj); err != nil {
		fmt.Println("JSON Unmarshal error", err)
		return
	}
	v, err := json.Marshal(yj)
	if err != nil {
		log.Fatal(err)
	}
	setHeadersHandler(w, r)
	fmt.Fprintf(w, "%s", v)
}

func setHeadersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Origin", r.Header.Get("Origin"))
	//domainによってAllowを可変する
	switch host := r.Header.Get("Origin"); host {
	case "http://10.0.1.6:8080", "http://10.0.1.6:8083", "http://localhost:8082", "http://10.0.1.6:3002", "http://10.0.1.6:8082'":
		w.Header().Set("Access-Control-Allow-Origin", host)
	default:
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//GETのみ許可
	//w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(http.StatusOK)
}
func RegisterHandlers() {
	//view api
	http.HandleFunc("/y", YoutubePage)
	http.HandleFunc("/j", Mock)

}
