package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qoharu/go-clean-ddd/config"
	"github.com/qoharu/go-clean-ddd/init"
	"github.com/qoharu/go-clean-ddd/internal/movie"
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
	init.Router(r, movieController)
}

func main() {
	_ = r.Run()
}
