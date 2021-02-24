package main

import (
	"ioc-container/cmd/controller"
	"ioc-container/pkg/http"
	"ioc-container/pkg/http/gin"
)

func main() {
	http.RegisterRouter(controller.TestController{})
	gin.RegisteGin()
}
