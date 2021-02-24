package gin

import (
	"sync"
)

var routers []Router
var once sync.Once

func initRouter() {
	once.Do(func() {
		routers = []Router{}
	})
}

// RegisterRouter 注册路由
func RegisterRouter(router Router) {
	initRouter()
	routers = append(routers, router)
}

// GetRouters 注册路由
func GetRouters() []Router {
	return routers
}
