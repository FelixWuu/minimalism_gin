package main

import (
	"github.com/FelixWuu/minimalism_gin/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.String(http.StatusOK, "路径输对了吗？")
	})

	router.LoadHTMLGlob("templates/*")

	router.GET("/rsp/string", response.StringResponse)
	router.GET("/rsp/json", response.JSONResponse)
	router.GET("rsp/html", response.HTMLResponse)
	router.GET("rsp/xml", response.XMLResponse)
	router.GET("rsp/yaml", response.YAMLResponse)

	router.Run()
}
