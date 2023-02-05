package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//type Article struct {
//	Title       string `json:"Title"`
//	Description string `json:"Description"`
//	Content     string `json:"Content"`
//}
//
//var Articles []Article

func main() {
	fmt.Println("Hello API!")
	//Articles = []Article{
	//	Article{Title: "How Politics Works", Description: "A way into politics", Content: "Article Content"},
	//	Article{Title: "How Cars Works", Description: "A way into cars", Content: "Article Content"},
	//	Article{Title: "How Stuff Works", Description: "A way into stuff", Content: "Article Content"},
	//	Article{Title: "How Food Works", Description: "A way into food", Content: "Article Content"},
	//}
	//fmt.Println(Articles)
	requestHandler()
}

func requestHandler() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	//mux.HandleFunc("/articles", returnAllArticles)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Endpoint hit: Home Page")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("View specific snippet"))
	fmt.Fprintf(w, "The id is %d", id)

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create Snippet"))
}

//func returnAllArticles(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Endpoint hit: Article Page")
//	json.NewEncoder(w).Encode(Articles)
//}
