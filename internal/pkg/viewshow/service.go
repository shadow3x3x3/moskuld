package viewshow

import (
	"errors"
	"log"
	"sync"
)

// Service represents a viewshow service
type Service interface {
	AddCinema(cinemas *Cinema) error
	AddMovie(movie *Movie) error
	GetCinemas() ([]*Cinema, error)
	GetMovies() ([]*Movie, error)

	GetSeats(sessionValue string)
}

type service struct {
	cinemas []*Cinema
	movies  []*Movie
}

// NewService returns a new viewshow service
func NewService() Service {
	return &service{}
}

// AddCinema adds a cinema to service
// If the cinema ID duplicates, returns an error.
func (s *service) AddCinema(cinema *Cinema) error {
	for _, c := range s.cinemas {
		if c.ID == cinema.ID {
			return errors.New("Duplicated Cinema")
		}
	}

	s.cinemas = append(s.cinemas, cinema)

	return nil
}

// AddMovie adds a movie to service
// If the movie ID duplicates, returns an error.
func (s *service) AddMovie(movie *Movie) error {
	for _, c := range s.movies {
		if c.ID == movie.ID {
			return errors.New("Duplicated Movie")
		}
	}

	s.movies = append(s.movies, movie)

	return nil
}

// GetCinemas return cinemas with four situations:
// 1. Cinemas and Movies both are null: Return all cinemas.
// 2. Cinemas is Null but Movies isn't null: Return cinemas include s.movies providing.
// 3. Cinemas isn't Null but Movies is null: Return s.cinemas at present.
// 4. Cinemas and Movies both aren't null: Return s.cinemas include s.movies providing.
func (s *service) GetCinemas() ([]*Cinema, error) {
	moviesNum := len(s.movies)
	cinemasNum := len(s.cinemas)

	allCinemas, err := getAllCinema()
	var (
		scopedCinemas []*Cinema
		scopedMovies  []*Movie
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

// GetMovies returns movies that provided by the s.cinemas
// TODO: Other situations.
func (s *service) GetMovies() ([]*Movie, error) {
	movies := []*Movie{}

	if len(s.cinemas) == 0 {
		return nil, errors.New("No Selected Cinema")
	}

	var wg sync.WaitGroup
	for _, c := range s.cinemas {
		wg.Add(1)

		go func(c *Cinema) {
			defer wg.Done()
			cMovies, err := getAllMovie(c.ID)
			if err != nil {
				log.Println("Can not get movie at", c.Name)
			}
			movies = append(movies, cMovies...)
		}(c)

	}
	wg.Wait()

	return movies, nil
}

func (s *service) GetSeats(sessionValue string) {
	getSeats(sessionValue)
}

func retriveCinemas(cinemas []*Cinema, movies []*Movie) ([]*Cinema, error) {
	cinemasWithMovie := []*Cinema{}
	for _, c := range cinemas {
		for _, m := range movies {
			if hasMovie(c.ID, m.ID) {
				cinemasWithMovie = append(cinemasWithMovie, c)
			}
		}
	}

	return cinemasWithMovie, nil
}

func hasMovie(cinemaID, movieID string) bool {
	movies, err := getAllMovie(cinemaID)
	if err != nil {
		return false
	}

	for _, m := range movies {
		if m.ID == movieID {
			return true
		}
	}

	return false
}

var _ Service = &service{}
