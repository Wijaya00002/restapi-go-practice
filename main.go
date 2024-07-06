package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

type Articles []Article

var articles = Articles{
	Article{Id: "541", Title: "Test Title", Description: "Test Description", Content: "Test Content"},
	Article{Id: "542", Title: "Test Title2", Description: "Test Description2", Content: "Test Content2"},
	Article{Id: "543", Title: "Test Title3", Description: "Test Description3", Content: "Test Content3"},
	Article{Id: "544", Title: "Test Title4", Description: "Test Description4", Content: "Test Content4"},
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func articleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["Id"]

	// fmt.Fprintf(w, "Key: "+key)

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}

	}
}

func getStringValue(values url.Values, key string) string {
	if v, ok := values[key]; ok && len(v) > 0 {
		return v[0]
	}
	return ""
}
func articleCreate(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	articles = append(articles, article)

	json.NewEncoder(w).Encode(articles)

	// fmt.Fprintf(w, "%+v", string(reqBody))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest() {

	Route := mux.NewRouter().StrictSlash(true)
	Route.HandleFunc("/", homePage)
	Route.HandleFunc("/articles/{Id}", articleDetail).Methods("GET")
	Route.HandleFunc("/articles", allArticles).Methods("GET")
	Route.HandleFunc("/articles", articleCreate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", Route))
}

func main() {
	fmt.Println("Rest API versi bapak hadi - Menggunakan Mux Routers")

	// newArticle := Article{Id: "545", Title: "New Test Title", Description: "New Test Description", Content: "New Test Content"}

	// articles = append(articles, newArticle)

	// articlesJSON, _ := json.MarshalIndent(articles, "", "  ")
	// fmt.Println(string(articlesJSON))
	handleRequest()
}
