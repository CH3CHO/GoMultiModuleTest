package main

import (
	"fmt"
	"github.com/CH3CHO/GoMultiModuleTest/core/service"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/pkg/container"
	"net/http"
)

var sayHelloService service.SayHelloService

func init() {
	ctn := &container.Container{}
	service.InitContainer(ctn)
	ctn.Make(&sayHelloService)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "You've reached test service.")
	})
	r.GET("/sayHello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "nil"
		}
		text := sayHelloService.SayHello(name)
		c.JSON(200, gin.H{
			"text": text,
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
