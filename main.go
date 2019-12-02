package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/my", myHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// indexHandler ... サインアップページ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("templates/signup.html"))
	if err := t.ExecuteTemplate(w, "signup.html", nil); err != nil {
		log.Fatal(err)
	}

}

// myHandler ... サインアップ後のマイページ
func myHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/my" {
		http.NotFound(w, r)
		return
	}

	t := template.Must(template.ParseFiles("templates/my.html"))
	if err := t.ExecuteTemplate(w, "my.html", nil); err != nil {
		log.Fatal(err)
	}
}
