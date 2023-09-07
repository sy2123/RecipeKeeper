package handlers

import (
	"html/template"
	"log"
    "net/http"
	"github.com/sy2123/RecipeKeeper/data"
)

func RecipePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/recipePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with recipe page file path")
	}
	id := r.URL.Path[len("/recipe/"):]
	var single data.Recipe

	for _, value := range data.AllRecipes {
		if value.Id == id {
			single = value
			break
		}
	}
	tmpl.Execute(w, single)
}