package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("User-Agent")
		log.Printf("%s - [%s %s %s] 200 [%s]", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, userAgent)
		w.Header().Set("Server", "Hello/v1")
		fmt.Fprintf(w, "Hello, World!\n")
	})

	log.Println("Starting Hello service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
