package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("ola mundo!"))
	http.ServeFile(w, r, "./templates/index.html")
}

func main() {

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
