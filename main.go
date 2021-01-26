package main

import (
	"github.com/qoharu/go-clean-ddd/config"
	controller2 "github.com/qoharu/go-clean-ddd/movie/controller"
	"github.com/qoharu/go-clean-ddd/movie/repository"
	"github.com/qoharu/go-clean-ddd/movie/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var r *gin.Engine

func init() {
	r = gin.Default()
	db := config.GetDBConnection()

	// Initialize Movie Service
	movieRepo := repository.NewMovieSQLRepository(db)
	movieUseCase := usecase.NewMovieUseCase(movieRepo)
	movieAPI := r.Group("/v1/movies")
	controller2.InitiateMovieHTTPController(movieAPI, movieUseCase)
}

func main() {
	r.Run()
}
