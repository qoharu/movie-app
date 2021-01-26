package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qoharu/go-clean-ddd/common/http/response"
	"github.com/qoharu/go-clean-ddd/movie"
	"net/http"
	"strconv"
)

type movieHTTPController struct {
	movieUseCase movie.UseCase
}

// InitiateMovieHTTPController ...
func InitiateMovieHTTPController(r *gin.RouterGroup, movieUseCase movie.UseCase) {
	controller := &movieHTTPController{
		movieUseCase: movieUseCase,
	}
	r.GET("/", controller.FindByTitle)
	r.GET("/:id", controller.FindByID)
	r.POST("/", controller.Add)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}

func (controller *movieHTTPController) FindByTitle(c *gin.Context) {
	movies, err := controller.movieUseCase.FindByTitle(c.Request.Context(), c.Query("title"))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, movies)
}

func (controller *movieHTTPController) FindByID(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.FindByID(c.Request.Context(), reqID)
	defaultMovie := movie.Movie{}
	if result == defaultMovie {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *movieHTTPController) Add(c *gin.Context) {
	var spec movie.Movie
	err := c.ShouldBindJSON(&spec)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.Add(c.Request.Context(), spec)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, result)
}

func (controller *movieHTTPController) Update(c *gin.Context) {
	var spec movie.Movie
	var err error
	err = c.ShouldBindJSON(&spec)

	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	spec.ID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.Update(c.Request.Context(), spec)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}

func (controller *movieHTTPController) Delete(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	err = controller.movieUseCase.Delete(c.Request.Context(), reqID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, nil)
}
