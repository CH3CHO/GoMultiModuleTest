package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
	"net/http"
	"github.com/CH3CHO/GoMultiModuleTest/core/service"
)

var ctn = &container.Container{}

func init() {
	service.InitContainer(ctn)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "You've reached test service.")
	})
	r.GET("/sayHello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"text": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
