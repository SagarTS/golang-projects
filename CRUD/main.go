package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json: "id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index,movie := range movies{
		if movie.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _,movie := range movies{
		if movie.ID == params["id"]{
			json.NewEncoder(w).Encode(movie)
			return
		}
	}	
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movies)
		}
	} 
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4749948", Title: "Avatar", Director: &Director{Firstname: "James", Lastname: "Cameroon"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4749949", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Println("Server starting at 8000 port")
	log.Fatal(http.ListenAndServe(":8000",r))
	
}