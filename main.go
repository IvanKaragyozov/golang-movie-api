package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Song struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"` // Duration in format "mm:ss"
}

var library []Song

func main() {
	library = []Song{
		{ID: uuid.New().String(), Title: "Song One", Artist: "Artist One", Album: "Album One", Year: 2001, Genre: "Rock", Duration: "3:45"},
		{ID: uuid.New().String(), Title: "Song Two", Artist: "Artist Two", Album: "Album Two", Year: 2005, Genre: "Pop", Duration: "4:05"},
		{ID: uuid.New().String(), Title: "Song Three", Artist: "Artist Three", Album: "Album Three", Year: 2010, Genre: "Jazz", Duration: "5:15"},
		{ID: uuid.New().String(), Title: "Song Four", Artist: "Artist Four", Album: "Album Four", Year: 2015, Genre: "Classical", Duration: "2:30"},
	}

	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/api/songs", getSongs).Methods("GET")
	router.HandleFunc("/api/songs/{id}", getSong).Methods("GET")
	router.HandleFunc("/api/songs", createSong).Methods("POST")
	router.HandleFunc("/api/songs/{id}", deleteSong).Methods("DELETE")
	router.HandleFunc("/api/songs/{id}", updateSong).Methods("PUT")

	fmt.Println("Server is running successfully on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return
	}
}

func getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, song := range library {
		if song.ID == params["id"] {
			json.NewEncoder(w).Encode(song)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newSong Song
	json.NewDecoder(r.Body).Decode(&newSong)
	newSong.ID = uuid.New().String()
	library = append(library, newSong)
	json.NewEncoder(w).Encode(newSong)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, song := range library {
		if song.ID == params["id"] {
			library = append(library[:i], library[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
}

func updateSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedSong Song
	json.NewDecoder(r.Body).Decode(&updatedSong)
	for i, song := range library {
		if song.ID == params["id"] {
			updatedSong.ID = song.ID
			library[i] = updatedSong
			json.NewEncoder(w).Encode(updatedSong)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
}
