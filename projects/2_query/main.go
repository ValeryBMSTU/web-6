package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		w.Write([]byte("Hello,"))
		w.Write([]byte(name))
		w.Write([]byte("!"))
	})

	http.ListenAndServe(":9000", nil)

}
