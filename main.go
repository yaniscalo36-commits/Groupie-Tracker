package main

import (
	"encoding/json"
	"log"
	"net/http"

	"groupie-tracker/api"
)

func main() {

	// Static files CSS/JS
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Static images (si tu mets des images locales dans ./image)
	imageFs := http.FileServer(http.Dir("./image"))
	http.Handle("/image/", http.StripPrefix("/image/", imageFs))

	// Pages HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.HandleFunc("/artists", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/artists.html")
	})

	// API locale qui fait proxy vers l’API distante
	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		artists, err := api.LoadArtistsFromAPI(
			"https://groupietrackers.herokuapp.com/api/artists",
		)
		if err != nil {
			http.Error(w, "Erreur API", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artists)
	})

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
