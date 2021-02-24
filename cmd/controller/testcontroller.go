package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (c TestController) GetPath() string {
	return "get"
}
func (c TestController) GetMethod() string {
	return "GET"
}

func (c TestController) GetHandler() func(*gin.Context) {
	return get
}

func get(this *gin.Context) {
	log.Println(">>>> hello gin start <<<<")
	this.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}
