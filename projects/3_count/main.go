package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

var count int

type req struct {
	Count string `json:"count"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data, _ := io.ReadAll(r.Body)
		var rq req
		json.Unmarshal(data, &rq)
		i, err := strconv.Atoi(rq.Count)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		count += i

	}
	if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(count)))
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", nil)
}
