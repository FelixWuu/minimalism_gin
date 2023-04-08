package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello NutCat!")
	})

	router.Run(":8080")
	// another way to start
	//http.ListenAndServe(":8080", router)
}
