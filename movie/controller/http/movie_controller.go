package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qoharu/go-clean-ddd/domain"
)

type movieController struct {
	movieUseCase domain.MovieUseCase
}

// NewMovieController ...
func NewMovieController(r *gin.RouterGroup, movieUseCase domain.MovieUseCase) {
	controller := &movieController{
		movieUseCase: movieUseCase,
	}
	r.GET("/movies", controller.FindByTitle)
	r.GET("/movies/:id", controller.FindByID)
	r.POST("/movies", controller.Add)
	r.PUT("/movies/:id", controller.Update)
	r.DELETE("/movies/:id", controller.Delete)
}

func (controller *movieController) FindByTitle(c *gin.Context) {
	movies, err := controller.movieUseCase.FindByTitle(c.Request.Context(), c.Query("title"))

	success := true
	errorMessage := ""
	if err != nil {
		success = false
		errorMessage = err.Error()
	
	}

	c.JSON(200, gin.H{
		"data":    movies,
		"success": success,
		"error": errorMessage,
	})
}

func (controller *movieController) FindByID(c *gin.Context) {
	reqID, _ := strconv.Atoi(c.Param("id"))
	movie, err := controller.movieUseCase.FindByID(c.Request.Context(), reqID)

	success := true
	errorMessage := ""
	if err != nil {
		success = false
		errorMessage = err.Error()
	}

	c.JSON(200, gin.H{
		"data":    movie,
		"success": success,
		"error": errorMessage,
	})
}

func (controller *movieController) Add(c *gin.Context) {
	var movie domain.Movie
	c.ShouldBindJSON(&movie)

	result, err := controller.movieUseCase.Add(c.Request.Context(), movie)

	success := true
	errorMessage := ""
	if err != nil {
		success = false
		errorMessage = err.Error()
	}

	c.JSON(200, gin.H{
		"data":    result,
		"success": success,
		"error": errorMessage,
	})
}

func (controller *movieController) Update(c *gin.Context) {
	var movie domain.Movie
	c.ShouldBindJSON(&movie)
	movie.ID, _ = strconv.Atoi(c.Param("id"))

	result, err := controller.movieUseCase.Update(c.Request.Context(), movie)

	success := true
	errorMessage := ""
	if err != nil {
		success = false
		errorMessage = err.Error()
	}

	c.JSON(200, gin.H{
		"data":    result,
		"success": success,
		"error": errorMessage,
	})
}

func (controller *movieController) Delete(c *gin.Context) {
	reqID, _ := strconv.Atoi(c.Param("id"))
	err := controller.movieUseCase.Delete(c.Request.Context(), reqID)

	success := true
	errorMessage := ""
	if err != nil {
		success = false
		errorMessage = err.Error()
	}

	c.JSON(200, gin.H{
		"data":    nil,
		"success": success,
		"error": errorMessage,
	})
}
