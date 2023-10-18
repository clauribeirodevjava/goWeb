package main

import (
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id     int
	Tittle string
	Body   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		post := Post{Id: 1, Tittle: "Unamed Post", Body: "No Content"}
		if title := r.FormValue("title"); title != "" {
			post.Tittle = title
		}

		t := template.Must(template.ParseFiles("templates/index.html"))
		if err := t.ExecuteTemplate(w, "index.html", post); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	fmt.Println(http.ListenAndServe(":8080", nil))
}
