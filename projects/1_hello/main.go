package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, web!"))
	})
	http.ListenAndServe(":8080", nil)
}
