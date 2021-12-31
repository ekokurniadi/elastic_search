package main

import (
	"context"
	"elastic_go/mapping"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("client is successfully connected")

	ctx := context.Background()
	isExist, err := client.IndexExists("belajar").Do(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	if !isExist {
		_, err = client.CreateIndex("belajar").Body(mapping.Mapping).Do(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Index belum ada")
	}

	api.GET("")
	router.Run()

}
