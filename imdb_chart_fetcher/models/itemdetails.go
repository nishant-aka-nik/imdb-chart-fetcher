package models

type MovieAttributes struct {
	Title              string  `json:"title"`
	Movie_release_year int     `json:"movie_release_year"`
	Imdb_rating        float64 `json:"imdb_rating"`
	Summary            string  `json:"summary"`
	Duration           string  `json:"duration"`
	Genre              string  `json:"genre"`
}
