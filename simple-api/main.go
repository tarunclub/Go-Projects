package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for movies
type Movie struct {
	MovieId     string    `json:"movie_id"`
	MovieName   string    `json:"movie_name"`
	TicketPrice int       `json:"price"`
	Director    *Director `json:"director"`
}

type Director struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var movies []Movie

// Middlewares
func (m *Movie) IsEmpty() bool {
	return m.MovieName == ""
}

func main() {
	r := mux.NewRouter()

	// seeding
	movies = append(movies, Movie{MovieId: "1", MovieName: "Test", TicketPrice: 188, Director: &Director{Fullname: "TestName", Website: "Test website"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movie", addOneMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getOneMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", deleteOneMovie).Methods("DELETE")
	r.HandleFunc("/movie/{id}", updateOneMovie).Methods("PUT")

	// Listen to Port
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to movies application API</h1>"))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, val := range movies {
		if val.MovieId == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found")
}

func addOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// validate - empty data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data")
	}

	// validate - {}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if movie.IsEmpty() {
		json.NewEncoder(w).Encode("Please add some data")
		return
	}

	// Generate a random string of length 10.
	randString := make([]byte, 10)
	for i := range randString {
		randString[i] = byte(rand.Intn(256))
	}

	// Convert the random string to a string.
	id := string(randString)

	movie.MovieId = id

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, val := range movies {
		if val.MovieId == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.MovieId = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, val := range movies {
		if val.MovieId == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			return
		}
	}
	json.NewEncoder(w).Encode("Not able to find the id")
}
