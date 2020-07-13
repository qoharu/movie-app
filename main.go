package main

import (
	"github.com/qoharu/go-clean-ddd/config"
	controller "github.com/qoharu/go-clean-ddd/movie/controller/http"
	"github.com/qoharu/go-clean-ddd/movie/repository"
	"github.com/qoharu/go-clean-ddd/movie/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


var r *gin.Engine

// Initialize Service
func init() {
	r = gin.Default()
	db := config.GetDBConnection()

	movieRepo := repository.NewMovieRepository(db)
	movieUseCase := usecase.NewMovieUseCase(movieRepo)
	movieAPI := r.Group("/v1")
	controller.NewMovieController(movieAPI, movieUseCase)
}

func main() {
	r.Run()
}
