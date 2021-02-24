package main

import (
	"github.com/machengling/go-ioc/cmd/controller"
	"github.com/machengling/go-ioc/pkg/http/gin"
)

func main() {
	gin.RegisterRouter(controller.TestController{})
	gin.RegisteGin(&gin.Config{IP: "localhost", Port: "8080"})
}
