package main

import (
	"go-ioc/cmd/controller"
	"go-ioc/pkg/http/gin"
)

func main() {
	gin.RegisterRouter(controller.TestController{})
	gin.RegisteGin(&gin.Config{IP: "localhost", Port: "8080"})
}
