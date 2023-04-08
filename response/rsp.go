package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StringResponse(c *gin.Context) {
	c.String(http.StatusOK, "Hello, NutCat! This is your string")
}

type JsonMsg struct {
	Name        string  `json:"name"`
	Score       float64 `json:"score"`
	Description string  `json:"description"`
}

func JSONResponse(c *gin.Context) {
	msg := JsonMsg{"NutCat", 99.9, "打工仔1号"}
	c.JSON(http.StatusOK, msg)
	c.JSON(http.StatusOK, gin.H{"Name": "FelixWuu", "Score": 100, "Description": "打工仔2号"})
}

func HTMLResponse(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		// Pass in the current time
		"time": time.Now().Format("2006-01-02 15:04:05"),
	})
}

func XMLResponse(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"user": "NutCat", "score": 100.00, "level": 3})
}

func YAMLResponse(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"user": "FelixWuu", "score": 99.99, "level": 10})
}
