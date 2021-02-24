package gin

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Config ...
type Config struct {
	IP   string
	Port string
}

// Router 处理器
type Router interface {
	GetHandler() map[string]func(*gin.Context)
}

// RegisteGin 注册gin
func RegisteGin(config *Config) {
	router := gin.Default()

	registeHandler(router)
	// 指定地址和端口号

	var ip string = "localhost"
	var port string = "9090"
	if config != nil && config.IP != "" {
		ip = config.IP
	}
	if config != nil && config.Port != "" {
		port = config.Port
	}
	router.Run(fmt.Sprintf("%s:%s", ip, port))
}

func registeHandler(engine *gin.Engine) {
	routerMap := GetRouters()
	for _, router := range routerMap {
		ginRouter := router.(Router)

		for path, handler := range ginRouter.GetHandler() {
			pathAndMethod := strings.Split(path, "@")

			switch pathAndMethod[0] {
			case "POST":
				engine.POST(pathAndMethod[1], handler)
			case "GET":
				engine.GET(pathAndMethod[1], handler)
			case "PUT":
				engine.PUT(pathAndMethod[1], handler)
			case "PATCH":
				engine.PATCH(pathAndMethod[1], handler)
			case "DELETE":
				engine.DELETE(pathAndMethod[1], handler)
			case "Any":
				engine.Any(pathAndMethod[1], handler)
			case "OPTIONS":
				engine.OPTIONS(pathAndMethod[1], handler)
			case "HEAD":
				engine.HEAD(pathAndMethod[1], handler)
			}
		}

	}
}
