package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/my", myHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// サインアップページ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Sign Up!")

}

// サインアップ後のマイページ
func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "My Page!")
}