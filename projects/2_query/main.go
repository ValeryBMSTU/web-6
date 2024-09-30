package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello," + r.URL.Query().Get("name") + "!"))
}
func main() {
	http.HandleFunc("/api/user/", handler)
	http.ListenAndServe(":9000", nil)
}
