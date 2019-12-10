package main

import (
	"log"
	"moskuld/internal/pkg/viewshow"
	"moskuld/pkg/cinema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	router.GET("/cinemas", getCinemas)
	router.GET("/cinema/:cinemaID")
	router.GET("/cinema/:cinemaID/movies", getMovies)
	router.GET("/cinema/:cinemaID/movies/:movieID")
}

func makeErrorResponse(c *gin.Context, code int, message string) {
	if code <= 200 {
		log.Println("Status <= 200 is not an error response")
	}
	c.JSON(code,
		gin.H{
			"status":  code,
			"message": message,
		})
}

type RespCinema struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func getCinemas(c *gin.Context) {
	vsService := viewshow.NewService()
	cinemas, err := vsService.GetCinemas()
	if err != nil {
		makeErrorResponse(c, http.StatusInternalServerError, err.Error())

		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get Cinmeas Success",
		"data":    cinemas,
	})

}

func getMovies(c *gin.Context) {
	cinemaID := c.Param("cinemaID")

	vsService := viewshow.NewService()
	err := vsService.AddCinema(&cinema.Cinema{ID: cinemaID})
	if err != nil {
		makeErrorResponse(c, http.StatusInternalServerError, err.Error())

		log.Println(err)
		return
	}

	movies, err := vsService.GetMovies()
	if err != nil {
		makeErrorResponse(c, http.StatusInternalServerError, err.Error())

		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get Movies Success",
		"data":    movies,
	})
}

func main() {
	r := gin.Default()

	registerRouter(r)

	r.Run()
}
