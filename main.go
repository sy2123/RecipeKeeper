package main

import (
	"net/http"
	"github.com/sy2123/RecipeKeeper/data"
	"github.com/sy2123/RecipeKeeper/handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.Println(data.AllRecipes)
	http.HandleFunc("/", handlers.HomePage)
	http.ListenAndServe(":8000", nil)
	
}