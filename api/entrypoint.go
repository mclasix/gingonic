package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
            "message": "Hello from Golang Gin Gonic",
        })
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w,r)
}