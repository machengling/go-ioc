package main

import (
	"ioc-container/cmd/controller"
	"ioc-container/pkg/http/gin"
)

func main() {
	gin.RegisterRouter(controller.TestController{})
	gin.RegisteGin(&gin.Config{IP: "localhost", Port: "8080"})
}
