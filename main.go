package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const staticPrefix string = "static/"

func messageHandler(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(message))
	}
}

func staticFileHandler(fileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(strings.Join([]string{staticPrefix, fileName}, ""))
		http.ServeFile(w, r, strings.Join([]string{staticPrefix, fileName}, ""))
	}
}

func printPathHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func main() {
	var addr = os.Getenv("ADDR") // type-inferred

	if len(addr) == 0 {
		addr = ":3000"
	}

	var mux = http.NewServeMux()

	mux.HandleFunc("/", messageHandler("Hello, world!"))
	mux.HandleFunc("/printpath/", printPathHandler)
	mux.HandleFunc("/go", staticFileHandler("go.html"))

	log.Printf("server is listening at %s", addr)
	var err = http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
