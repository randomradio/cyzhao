package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func addUlr(w http.ResponseWriter, r *http.Request) {

	url := r.FormValue("url")

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	w.Write(body)
}

func main() {
	fmt.Println("Start server")

	http.HandleFunc("/add_url", addUlr)

	http.ListenAndServe(":9001", nil)
}
