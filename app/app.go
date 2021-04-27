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
	Logline string `json:"Logline"`
	Content string `json:"Content"`
}

// var Articles []Article

func (app *App) SetupRouter() {
	app.Router.HandleFunc("/", homePage)
	app.Router.HandleFunc("/articles", app.createNewArticle).Methods("POST")
	app.Router.HandleFunc("/articles", app.returnAllArticles)
	app.Router.HandleFunc("/articles/{id}", app.deleteArticle).Methods("DELETE")
	app.Router.HandleFunc("/articles/{id}", app.updateArticle).Methods("PUT")
	app.Router.HandleFunc("/articles/{id}", app.returnSingleArticle)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func (app *App) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	var Articles []Article

	rows, err := app.Database.Query("select * from articles;")
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Logline, &article.Content)
		if err != nil {
			log.Fatal("Database SCAN failed")
		}
		Articles = append(Articles, article)
	}

	json.NewEncoder(w).Encode(Articles)
}

func (app *App) returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := app.Database.QueryRow("select * from articles where id = ?", id)

	var article Article
	err := row.Scan(&article.Id, &article.Title, &article.Logline, &article.Content)
	if err != nil {
		json.NewEncoder(w).Encode("No article found with that ID")
	} else {
		json.NewEncoder(w).Encode(article)
	}
}

func (app *App) createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	_, err := app.Database.Exec("INSERT INTO `articles` (id, title, logline, content) VALUES (?, ?, ?, ?)", article.Id, article.Title, article.Logline, article.Content)
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	json.NewEncoder(w).Encode(article)
}

func (app *App) deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := app.Database.Exec("delete from articles where id = ?", id)
	if err != nil {
		json.NewEncoder(w).Encode("Database DELETE failed")
	} else {
		json.NewEncoder(w).Encode("Article deleted")
	}
}

func (app *App) updateArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updated Article
	json.Unmarshal(reqBody, &updated)

	_, err := app.Database.Exec("update articles set id = ?, title = ?, logline = ?, content = ? where id = ?", updated.Id, updated.Title, updated.Logline, updated.Content, id)
	if err != nil {
		json.NewEncoder(w).Encode("Database UPDATE failed")
	} else {
		json.NewEncoder(w).Encode(updated)
	}
}
