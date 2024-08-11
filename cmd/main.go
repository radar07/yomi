package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/radar07/yomi/views"
)

func render(c *gin.Context, status int, template templ.Component) {
	c.Status(status)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
		component := views.Hello()
		component.Render(ctx, ctx.Writer)
	})

	router.Run(":8080")
}
