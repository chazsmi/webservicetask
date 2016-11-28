package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Reponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func sendReponse(
	res string,
	w http.ResponseWriter,
	code int,
) {
	if code == 500 {
		res = "Bad request"
	}
	r := Reponse{
		Code:   code,
		Result: res,
	}
	js, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(js)
}

func hello(w http.ResponseWriter, r *http.Request) {
	// Build the response struct
	name := r.FormValue("name")
	code := 200
	res := ""
	if len(name) == 0 {
		code = 422
	} else {
		res = fmt.Sprintf("%s %s", "hello", name)
	}

	sendReponse(res, w, code)
}

func main() {
	port := flag.String("port", "8000", "Server port")
	flag.Parse()

	http.HandleFunc("/", hello)

	// Starts the web server
	http.ListenAndServe(":"+*port, nil)
}
