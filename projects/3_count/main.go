package main

import (
	"net/http"
	"strconv"
)

func main() {
	var counter int = 0
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(strconv.Itoa(counter)))
		case http.MethodPost:
			r.ParseForm()
			st := r.Form.Get("count")
			i, wrong := strconv.Atoi(st)
			if wrong != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
			} else {
				counter = counter + i
			}
		}
	})

	http.ListenAndServe(":3333", nil)

}
