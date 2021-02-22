package ioc

import (
	"fmt"
	"ioc-container/pkg/utils"
	"sync"
)

// Container 容器
type Container struct {
	// components 实际组件存放的地方
	components map[string]*Component
	lock       sync.Mutex
}

// containerIns 全局的容器实例
var containerIns *Container
var initOnce sync.Once

// InitContainer 初始化容器
func InitContainer() {
	initOnce.Do(func() {
		container := Container{}
		container.components = map[string]*Component{}
		containerIns = &container
	})
}

// RegisteComponent 添加组件
func RegisteComponent(component *Component) (bool, error) {
	InitContainer()
	containerIns.lock.Lock()
	defer containerIns.lock.Unlock()
	fmt.Println("component.name", component.name)

	if _, ok := containerIns.components[component.name]; ok {
		panic("[ioc-container] component has same name:" + component.name)
	}
	containerIns.components[component.name] = component
	return true, nil
}

// GetComponent 获取组件
func GetComponent(component *Component) (*Component, error) {
	containerIns.lock.Lock()
	defer containerIns.lock.Unlock()
	strName := component.name
	fmt.Println("strName", strName)
	if component.name == "" {
		strName = utils.GetComponentName(component.obj)
	}
	if c, ok := containerIns.components[strName]; ok {
		return c, nil
	}
	return nil, nil
}
