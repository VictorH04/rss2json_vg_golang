package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func rss2json(w http.ResponseWriter,r *http.Request)  {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.vg.no/rss/feed?categories=sport")
	json.NewEncoder(w).Encode(feed.Items)


	// for _, item := range feed.Items {
	// 	fmt.Println(item.Title)
	// 	fmt.Println(item.Description)
	// 	fmt.Println(item.Content)
	// }

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/vg_json", rss2json)
    log.Fatal(http.ListenAndServe(":8822", myRouter))
}


func main() {
    handleRequests()
}