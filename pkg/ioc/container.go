package ioc

import (
	"ioc-container/pkg/utils"
	"sync"
)

// Container 容器
type Container struct {
	// components 实际组件存放的地方
	components map[string]AbstractComponent
	// controllers 控制器组件存放的地方
	lock sync.Mutex
}

// containerIns 全局的容器实例
var containerIns *Container
var initOnce sync.Once

// InitContainer 初始化容器
func InitContainer() {
	initOnce.Do(func() {
		container := Container{}
		container.components = map[string]AbstractComponent{}
		containerIns = &container
	})
}

// RegisteComponent 添加组件
func RegisteComponent(component AbstractComponent) (bool, error) {
	InitContainer()
	containerIns.lock.Lock()
	defer containerIns.lock.Unlock()

	if _, ok := containerIns.components[component.GetName()]; ok {
		panic("[ioc-container] component has same name:" + component.GetName())
	}
	containerIns.components[component.GetName()] = component
	return true, nil
}

// GetComponent 获取组件
func GetComponent(component AbstractComponent) (AbstractComponent, error) {
	containerIns.lock.Lock()
	defer containerIns.lock.Unlock()
	strName := component.GetName()
	if component.GetName() == "" {
		strName = utils.GetComponentName(component.GetObj())
	}
	if c, ok := containerIns.components[strName]; ok {
		return c, nil
	}
	return nil, nil
}
