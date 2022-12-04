package structs

import (
	"github.com/google/uuid"
	"time"
)

type Movie struct {
	ID            int
	Title         string
	Author        string
	AuthorID      int
	PublishedData time.Time
}

type Author struct {
	ID        int
	Name      string
	Janr      string
	Movies    []int
	BirthData time.Time
	DeathData time.Time
}

type Genres struct {
	ID        uuid.UUID `json:"id"`
	GenreName string    `json:"genrename"`
}

type MoiveGenres struct {
	MoiveID uuid.UUID `json:"moiveid"`
	GenreID uuid.UUID `json:"genreID"`
}

type Moives struct {
	ID          uuid.UUID `json:"id"`
	DirectorID  uuid.UUID `json:"directorid"`
	GenreID     uuid.UUID `json:"genreid"`
	Title       string    `json:"title"`
	ReleaseYear int       `json:"releaseyear"`
	Rating      int       `json:"rating"`
	Plot        int       `json:"plot"`
	MovieLength time.Time `json:"movielength"`
}

type MovieActor struct {
	MovieID uuid.UUID `json:"movieid"`
	ActorID uuid.UUID `json:"actorid"`
}

type Actors struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Nationality string    `json:"nationality"`
	Birth       time.Time `json:"birth"`
}

type Directors struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Nationality string    `json:"nationality"`
	Birth       time.Time `json:"birth"`
}
