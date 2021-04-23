package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

type Article struct {
	Id      string `json."Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

func (app *App) SetupRouter() {
	app.Router.HandleFunc("/", homePage)
	app.Router.HandleFunc("/articles", createNewArticle).Methods("POST")
	app.Router.HandleFunc("/articles", returnAllArticles)
	app.Router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	app.Router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	app.Router.HandleFunc("/articles/{id}", returnSingleArticle)
	app.Router.HandleFunc("/test", app.testFunction).Methods("POST")
}

func (app *App) testFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('yourname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You called a thing!")
	w.WriteHeader(http.StatusOK)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	
	fmt.Println("Endpoint Hit: updateArticle")

	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updated Article
	json.Unmarshal(reqBody, &updated)

	for i, article := range Articles {
		if article.Id == id {
			article.Title = updated.Title
			article.Desc = updated.Desc
			article.Content = updated.Content

			Articles[i] = article

			json.NewEncoder(w).Encode(article)
		}
	}
}
