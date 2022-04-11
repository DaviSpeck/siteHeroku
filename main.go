package main

import (
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
