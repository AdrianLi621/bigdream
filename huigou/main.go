package main

import (
	"github.com/gin-gonic/gin"
	"io"
)

func main(){
	g:=gin.Default()
	g.GET("/", func(context *gin.Context) {
		io.WriteString(context.Writer,"success")
	})
	g.Run(":8899")
}