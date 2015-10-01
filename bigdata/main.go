package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/bigdata/json/", dash1)
	http.ListenAndServe(":9003", nil)
}
func dash1(w http.ResponseWriter, rw *http.Request) {
	if origin := rw.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if rw.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	http.ServeFile(w, rw, rw.URL.Path[1:])
	fmt.Println(rw.URL.Path[1:])
}
