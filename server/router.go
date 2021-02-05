package server

import (
	"github.com/gin-gonic/gin"
	"github.com/qoharu/movie-app/internal/movie"
)

func registerMovieRoute(r *gin.Engine, movieController movie.HTTPController) {
	movieRouter := r.Group("/v1/movies")
	movieRouter.GET("/", movieController.FindByTitle)
	movieRouter.GET("/:id", movieController.FindByID)
	movieRouter.POST("/", movieController.Add)
	movieRouter.PUT("/:id", movieController.Update)
	movieRouter.DELETE("/:id", movieController.Delete)
}
