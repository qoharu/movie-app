package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qoharu/movie-app/config"
	"github.com/qoharu/movie-app/internal/movie"
	"github.com/qoharu/movie-app/server"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	db := config.GetDBConnection()

	// Initialize Movie Service
	movieRepo := movie.NewRepository(db)
	movieUseCase := movie.NewUseCase(movieRepo)
	movieController := movie.NewHTTPController(movieUseCase)

	// Start API
	server.StartRouter(r, movieController)
}

func main() {
	_ = r.Run()
}
