package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/qoharu/go-clean-ddd/domain"
)

// DBCon ... Database Connection Instance
var dbCon *gorm.DB

// InitDB ... function to initialize database
func InitDB() {
	var err error

	var dbHost string = "localhost"
	var dbPort string = "5433"
	var dbUser string = "postgres"
	var dbPass string = "rahasia123"
	var dbName string = "movie"

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	//connect to postgres database
	dbCon, err = gorm.Open("postgres", dbString)

	//migrate database
	dbCon.AutoMigrate(&domain.Movie{})

	if err != nil {
		panic(err)
	}

}

//GetDBConnection ...
func GetDBConnection() *gorm.DB {
	InitDB()

	return dbCon
}
