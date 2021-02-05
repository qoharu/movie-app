package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qoharu/movie-app/config"
	"github.com/qoharu/movie-app/internal/movie"
)

var db *gorm.DB

func RegisterAPIService(e *gin.Engine) {
	db = config.GetDBConnection()

	registerMovieAPIService(e)
}

func registerMovieAPIService(r *gin.Engine) {
	// Initialize Movie Service
	movieRepo := movie.NewRepository(db)
	movieUseCase := movie.NewUseCase(movieRepo)
	movieController := movie.NewHTTPController(movieUseCase)

	// Start API
	registerMovieRoute(r, movieController)
}
