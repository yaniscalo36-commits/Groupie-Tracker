package main

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/api"
)

func main() {

	// Fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// Page accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Page artistes
	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		artists, err := api.LoadArtists()
		if err != nil {
			http.Error(w, "Erreur chargement artistes", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/artists.html"))
		tmpl.Execute(w, artists)
	})

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
