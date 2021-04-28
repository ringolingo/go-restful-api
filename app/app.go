package app

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ringolingo/go-restful-api/app/article"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.HandleFunc("/articles", app.createNewArticle).Methods("POST")
	app.Router.HandleFunc("/articles", app.returnAllArticles)
	app.Router.HandleFunc("/articles/{id}", app.deleteArticle).Methods("DELETE")
	app.Router.HandleFunc("/articles/{id}", app.updateArticle).Methods("PUT")
	app.Router.HandleFunc("/articles/{id}", app.returnSingleArticle)
}

func (app *App) returnAllArticles(w http.ResponseWriter, r *http.Request) {
	var Articles []article.Article

	rows, err := app.Database.Query("select * from articles;")
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	for rows.Next() {
		var a article.Article
		err := rows.Scan(&a.Id, &a.Title, &a.Logline, &a.Content)
		if err != nil {
			log.Fatal("Database SCAN failed")
		}
		Articles = append(Articles, a)
	}

	json.NewEncoder(w).Encode(Articles)
}

func (app *App) returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := app.Database.QueryRow("select * from articles where id = ?", id)

	var a article.Article
	err := row.Scan(&a.Id, &a.Title, &a.Logline, &a.Content)
	if err != nil {
		json.NewEncoder(w).Encode("No article found with that ID")
	} else {
		json.NewEncoder(w).Encode(a)
	}
}

func (app *App) createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var a article.Article
	json.Unmarshal(reqBody, &a)
	a.Title = article.FormatTitle(a.Title)
	a.Logline = article.FormatLogline(a.Logline)

	result, err := app.Database.Exec("INSERT INTO `articles` (title, logline, content) VALUES (?, ?, ?)", a.Title, a.Logline, a.Content)
	if err != nil {
		json.NewEncoder(w).Encode("Database INSERT failed")
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			json.NewEncoder(w).Encode("Article could not be added")
		} else {
			a.Id = id
			json.NewEncoder(w).Encode(a)
		}
	}
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

	var updated article.Article
	json.Unmarshal(reqBody, &updated)
	updated.Title = article.FormatTitle(updated.Title)
	updated.Logline = article.FormatLogline(updated.Logline)

	_, err := app.Database.Exec("update articles set title = ?, logline = ?, content = ? where id = ?", updated.Title, updated.Logline, updated.Content, id)
	if err != nil {
		json.NewEncoder(w).Encode("Database UPDATE failed")
	} else {
		updated.Id, _ = strconv.ParseInt(id, 10, 64)
		json.NewEncoder(w).Encode(updated)
	}
}
