package main

import (
	"log"
	"net/http"
	"github.com/sy2123/RecipeKeeper/data"
	"github.com/sy2123/RecipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.Println(data.AllRecipes)

	server := http.NewServeMux()

	style := http.FileServer(http.Dir("frontend/css"))

	server.Handle("/frontend/css/", http.StripPrefix("/frontend/css/", style))

	server.HandleFunc("/", handlers.HomePage)
	server.HandleFunc("/recipe/", handlers.RecipePage)

	log.Fatal(http.ListenAndServe(":8000", server))
	    
}