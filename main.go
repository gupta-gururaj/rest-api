//REST API implementation

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Article struct (Model)
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Init ariticles var as a slice Article struct
var articles []Article

// Get all articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func getSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets all parameteres from URL
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Update existing article
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	fmt.Fprintf(w, "Put Method")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...) //delete current book
			var newArticle Article
			_ = json.NewDecoder(r.Body).Decode(&newArticle)
			newArticle.ID = params["id"]
			articles = append(articles, newArticle)
			json.NewEncoder(w).Encode(articles)
			return
		}
	}
}

// Main function
func main() {

	//initialize base value
	articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Greetings", Content: "Have a Nice Day !!!"},
		Article{ID: "2", Title: "Bye", Desc: "Goodbye", Content: "Nice to meet you !!!"},
	}

	//Init router
	myRouter := mux.NewRouter()

	// Route handles & endpoints
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")        //get all data
	myRouter.HandleFunc("/articles/{id}", getSingleArticle).Methods("GET") //get single data
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")    //update single existing data

	// Start server
	http.ListenAndServe(":9050", myRouter)

	fmt.Println("Server - http://localhost:9050/")
}
