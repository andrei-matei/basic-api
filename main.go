package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var Articles []Article

func main() {
	fmt.Println("Hello API!")
	Articles = []Article{
		Article{Title: "How Politics Works", Description: "A way into politics", Content: "Article Content"},
		Article{Title: "How Cars Works", Description: "A way into cars", Content: "Article Content"},
		Article{Title: "How Stuff Works", Description: "A way into stuff", Content: "Article Content"},
		Article{Title: "How Food Works", Description: "A way into food", Content: "Article Content"},
	}
	fmt.Println(Articles)
	requestHandler()
}

func requestHandler() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", articles)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Endpoint hit: Home Page")
}

func articles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: Article Page")
	json.NewEncoder(w).Encode(Articles)
}
