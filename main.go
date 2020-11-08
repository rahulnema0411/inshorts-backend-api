package main

import (
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID                primitive.ObjectID
	Title             string
	Subtitle          string
	Content           string
	Created_TimeStamp time.Time
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/article", GetArticleEndPoint)
	http.ListenAndServe(":27017", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Just saying hello, %s!", r.URL.Path[1:])
}

func GetArticleEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The article end point, %s!", r.URL.Path[1:])
}
