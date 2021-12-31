package main

import (
	"elastic_go/handler"
	"elastic_go/repository"
	"elastic_go/service"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/olivere/elastic"
)

func main() {

	env := godotenv.Load()
	if env != nil {
		log.Fatal("cannot load the .env file")
	}

	host := os.Getenv("HOST")
	userHost := os.Getenv("USER_HOST")
	userPass := os.Getenv("USER_PASS")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePort := os.Getenv("DATABASE_PORT")

	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("client is successfully connected")

	tweetRepository := repository.NewTweetRepository(db)
	tweetService := service.NewTweetService(tweetRepository)
	tweetHandler := handler.NewTweetHandler(tweetService, client)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")
	api.POST("/tweet", tweetHandler.Save)
	router.Run()

}
