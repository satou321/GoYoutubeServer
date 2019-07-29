package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func init() {
	RegisterHandlers()
}

type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonError)
}
func StartWebServer(port string) error {
	fmt.Println(port)
	addr := ":80"
	log.Println("[infooo] http server listening", addr)
	return http.ListenAndServe(addr, nil)
}
