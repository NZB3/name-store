package main

import (
	"html/template"
	"log"
	"name-storage/database"
	"name-storage/database/models"
	"net/http"

	"github.com/lib/pq"
)

func home_page(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./page/index.html")
		if err != nil {
			log.Fatal("error parse template")
		}

		if r.Method == "POST" {
			r.ParseForm()
			name := r.Form.Get("name")
			email := r.Form.Get("email")
			user := models.NewUser(name, email)
			log.Println("Created: ", user)
			db.AddUser(user)
		}

		users, err := db.AllUsers()
		log.Println(users)
		if err != nil {
			log.Fatalf("error AllUsers %s", err)
		}
		tmpl.Execute(w, users)
	}
}

func main() {
	_ = pq.Array
	db, err := database.New("user=postgres dbname=name_db sslmode=disable")
	if err != nil {
		log.Fatalf("error connect db: %s", err)
	}
	http.HandleFunc("/", home_page(db))
	http.ListenAndServe(":8000", nil)
}
