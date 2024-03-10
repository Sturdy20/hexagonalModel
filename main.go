package main

import (
	"HexagonalModel/modules/servers/project1"
	"HexagonalModel/pkg/databases/postgresql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := postgresql.Postgresql()
	defer db.Close()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	router.Use(cors.New(config))

	project1.Project1(router, db)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
