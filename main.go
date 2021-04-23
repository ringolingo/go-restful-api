package main

import (
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/ringolingo/go-restful-api/internal/app"
	"github.com/ringolingo/go-restful-api/internal/db"	
)

// func handleRequests() {
// 	// myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
// 	myRouter.HandleFunc("/articles", returnAllArticles)
// 	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
// 	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
// 	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homepage")
// }

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnSingleArticle")
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	for _, article := range Articles {
// 		if article.Id == key {
// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: createNewArticle")
// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var article Article
// 	json.Unmarshal(reqBody, &article)
// 	Articles = append(Articles, article)
// 	json.NewEncoder(w).Encode(article)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 		}
// 	}
// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: updateArticle")

// 	id := mux.Vars(r)["id"]
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var updated Article
// 	json.Unmarshal(reqBody, &updated)

// 	for i, article := range Articles {
// 		if article.Id == id {
// 			article.Title = updated.Title
// 			article.Desc = updated.Desc
// 			article.Content = updated.Content

// 			Articles[i] = article

// 			json.NewEncoder(w).Encode(article)
// 		}
// 	}
// }

// type Article struct {
// 	Id      string `json."Id"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"Desc"`
// 	Content string `json:"Content"`
// }

// var Articles []Article

// func main() {
// 	Articles = []Article{
// 		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}
// 	handleRequests()
// }

func main() {
	
}
