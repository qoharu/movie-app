package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qoharu/movie-app/server"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	server.RegisterAPIService(r)
}

func main() {
	_ = r.Run()
}
