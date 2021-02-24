package gin

import (
	"ioc-container/pkg/http"

	"github.com/gin-gonic/gin"
)

// GinHandlerFunc 处理器
type GinHandlerFunc interface {
	GetHandler() func(*gin.Context)
}

// RegisteGin 注册gin
func RegisteGin() {
	router := gin.Default()

	registeHandler(router)
	// 指定地址和端口号
	router.Run("localhost:9090")
}

func registeHandler(engine *gin.Engine) {
	routerMap := http.GetRouters()
	for _, item := range routerMap {
		ginHandler := item.(GinHandlerFunc)
		switch item.GetMethod() {
		case "POST":
			engine.POST(item.GetPath(), ginHandler.GetHandler())
		case "GET":
			engine.GET(item.GetPath(), ginHandler.GetHandler())
		case "PUT":
			engine.PUT(item.GetPath(), ginHandler.GetHandler())
		case "PATCH":
			engine.PATCH(item.GetPath(), ginHandler.GetHandler())
		case "DELETE":
			engine.DELETE(item.GetPath(), ginHandler.GetHandler())
		case "Any":
			engine.Any(item.GetPath(), ginHandler.GetHandler())
		case "OPTIONS":
			engine.OPTIONS(item.GetPath(), ginHandler.GetHandler())
		case "HEAD":
			engine.HEAD(item.GetPath(), ginHandler.GetHandler())
		}
	}
}
