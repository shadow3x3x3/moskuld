package service

import (
	"errors"
	"moskuld/internal/pkg/cinema"
	"moskuld/internal/pkg/movie"
)

type Service interface {
	AddCinema(cinemas *cinema.Cinema) error
	GetCinemas() ([]*cinema.Cinema, error)
	AddMovie(movie *movie.Movie) error
}

type service struct {
	cinemas []*cinema.Cinema
	movies  []*movie.Movie
}

func NewService() Service {
	return &service{}
}

// GetCinemas return cinemas with four situations
// 1. Cinemas and Movies both are null: Return all cinemas
// 2. Cinemas is Null but Movies isn't null: Return cinemas include s.movies providing
// 3. Cinemas isn't Null but Movies is null: Return s.cinemas at present
// 4. Cinemas and Movies both aren't null: Return s.cinemas include s.movies providing
func (s *service) GetCinemas() ([]*cinema.Cinema, error) {
	moviesNum := len(s.movies)
	cinemasNum := len(s.cinemas)

	allCinemas, err := cinema.GetAll()
	var (
		scopedCinemas []*cinema.Cinema
		scopedMovies  []*movie.Movie
	)

	switch {
	case cinemasNum == 0 && moviesNum == 0:
		return allCinemas, err
	case cinemasNum == 0 && moviesNum != 0:
		scopedCinemas = allCinemas
		scopedMovies = s.movies
	case cinemasNum != 0 && moviesNum == 0:
		return s.cinemas, nil
	case cinemasNum != 0 && moviesNum != 0:
		scopedCinemas = s.cinemas
		scopedMovies = s.movies

	}

	return retriveCinemas(scopedCinemas, scopedMovies)
}

func (s *service) AddCinema(cinema *cinema.Cinema) error {
	for _, c := range s.cinemas {
		if c.ID == cinema.ID {
			return errors.New("Duplicated Cinema")
		}
	}

	s.cinemas = append(s.cinemas, cinema)

	return nil
}

func (s *service) AddMovie(movie *movie.Movie) error {
	for _, c := range s.movies {
		if c.ID == movie.ID {
			return errors.New("Duplicated Movie")
		}
	}

	s.movies = append(s.movies, movie)

	return nil
}

func retriveCinemas(cinemas []*cinema.Cinema, movies []*movie.Movie) ([]*cinema.Cinema, error) {
	cinemasWithMovie := []*cinema.Cinema{}
	for _, c := range cinemas {
		for _, m := range movies {
			if cinema.HaveMovie(c.ID, m.ID) {
				cinemasWithMovie = append(cinemasWithMovie, c)
			}
		}
	}

	return cinemasWithMovie, nil
}

var _ Service = &service{}
