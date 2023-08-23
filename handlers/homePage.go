package handlers

import(
	"net/http"
	"fmt"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintf(w, "Hello C4W!")
}