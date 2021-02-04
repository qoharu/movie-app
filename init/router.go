package init

import (
	"github.com/gin-gonic/gin"
	"github.com/qoharu/go-clean-ddd/internal/movie"
)

// Router ...
func Router(r *gin.Engine, movieController movie.HTTPController) {

	movieRouter := r.Group("/v1/movies")
	movieRouter.GET("/", movieController.FindByTitle)
	movieRouter.GET("/:id", movieController.FindByID)
	movieRouter.POST("/", movieController.Add)
	movieRouter.PUT("/:id", movieController.Update)
	movieRouter.DELETE("/:id", movieController.Delete)

}
