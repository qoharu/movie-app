package config

import (
	"fmt"
	"github.com/qoharu/go-clean-ddd/internal/movie"

	"github.com/jinzhu/gorm"
)

// DBCon ... Database Connection Instance
var dbCon *gorm.DB

// InitDB ... function to initialize database
func InitDB() {
	var err error

	var dbHost = "movie-postgres"
	var dbPort = "5432"
	var dbUser = "movie"
	var dbPass = "rahasia123"
	var dbName = "movie"

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	//connect to postgres database
	dbCon, err = gorm.Open("postgres", dbString)

	//migrate database
	dbCon.AutoMigrate(&movie.Movie{})

	if err != nil {
		panic(err)
	}

}

//GetDBConnection ...
func GetDBConnection() *gorm.DB {
	InitDB()

	return dbCon
}
