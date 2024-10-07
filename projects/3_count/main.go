package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var count int

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(count)))
		return
	} else if r.Method == "POST" {
		r.ParseForm()
		numS := r.Form.Get("count")
		num, err := strconv.Atoi(numS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}
		count += num
	}
}

func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
