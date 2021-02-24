package http

import (
	"sync"
)

// HandlerFunc 处理器
type HandlerFunc interface {
	GetPath() string
	GetMethod() string
}

var routers map[string]HandlerFunc
var once sync.Once

func initRouter() {
	once.Do(func() {
		routers = map[string]HandlerFunc{}
	})
}

// RegisterRouter 注册路由
func RegisterRouter(handle HandlerFunc) {
	initRouter()
	if _, ok := routers[handle.GetPath()+":"+handle.GetMethod()]; ok {
		panic("RegisterRouter handle path is exist")
	}
	routers[handle.GetPath()+":"+handle.GetMethod()] = handle
}

// GetRouters 注册路由
func GetRouters() map[string]HandlerFunc {
	return routers
}
