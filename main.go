package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"Title"`
	Subtitle string `json:"desc"`
	Content  string `json:"content"`
}

var Articles []Article

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(response).Encode(Articles)
}

func createArticle(response http.ResponseWriter, request *http.Request) {
	fmt.Println("methods", request.Method)
	if request.Method == "GET" {
		t, _ := template.ParseFiles("createArticle.html")
		t.Execute(response, nil)
	} else {
		request.ParseForm()
		// logic part of log in
		fmt.Println("ID:", request.Form["ID"])
		fmt.Println("Title:", request.Form["Title"])
		fmt.Println("Subtitle:", request.Form["Subtitle"])
		fmt.Println("Content:", request.Form["Content"])
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/articles", createArticle)
	http.ListenAndServe(":27017", nil)
}

func handleResponse() {

}

func main() {
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
