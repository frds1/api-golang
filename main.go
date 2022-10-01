package main

import (
	"desafio/interfaces/search"
	"desafio/interfaces/swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	r := router.Group("/")
	search.Router(r)
	swagger.Router(r)

	router.Run(":8080")
}
