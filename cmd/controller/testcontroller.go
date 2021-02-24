package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (c TestController) GetHandler() map[string]func(*gin.Context) {
	return map[string]func(*gin.Context){
		"GET@get": get,
	}
}

func get(this *gin.Context) {
	log.Println(">>>> hello gin start <<<<")
	this.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}
