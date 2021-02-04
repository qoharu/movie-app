package movie

import (
	"github.com/gin-gonic/gin"
	"github.com/qoharu/movie-app/internal/pkg/common/http/response"
	"net/http"
	"strconv"
)

type HTTPController interface {
	FindByTitle(c *gin.Context)
	FindByID(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// NewHTTPController ...
func NewHTTPController(movieUseCase UseCase) HTTPController {
	return &httpController{
		movieUseCase: movieUseCase,
	}
}

type httpController struct {
	movieUseCase UseCase
}

func (controller *httpController) FindByTitle(c *gin.Context) {
	movies, err := controller.movieUseCase.FindByTitle(c.Request.Context(), c.Query("title"))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, movies)
}

func (controller *httpController) FindByID(c *gin.Context) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := controller.movieUseCase.FindByID(c.Request.Context(), reqID)
	defaultMovie := Movie{}
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

func (controller *httpController) Add(c *gin.Context) {
	var spec Movie
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

func (controller *httpController) Update(c *gin.Context) {
	var spec Movie
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

func (controller *httpController) Delete(c *gin.Context) {
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
