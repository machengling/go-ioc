package ioc

import (
	"encoding/json"
	"fmt"
	"ioc-container/pkg/utils"
)

// AbstractComponent 组件抽象化
type AbstractComponent interface {
	SetName(string)
	GetName() string
	GetObj() interface{}
	SetObj(interface{})
	ConvToStruct(interface{})
}

// Component 组件
type Component struct {
	name string
	obj  interface{}
}

// NewComponent 创建一个组件
func NewComponent(obj interface{}) AbstractComponent {
	component := Component{}

	// 获取对象的结构体名称，作为组件的注册名称
	strName := utils.GetComponentName(obj)
	component.name = strName
	component.obj = obj
	return component
}

// ToString 打印组件信息
func (comp Component) ToString() string {
	return fmt.Sprintf("Component(name=%v;)", comp.name)
}

// SetName 设置组件名称，默认是使用组件结构体名称，或结构定义的ConponentName返回的字符串
func (comp Component) SetName(name string) {
	comp.name = name
}

// GetName 获取组件名称
func (comp Component) GetName() string {
	return comp.name
}

// SetObj 设置组件实体
func (comp Component) SetObj(obj interface{}) {
	comp.obj = obj
}

// GetObj 获取组件实体
func (comp Component) GetObj() (obj interface{}) {
	return comp.obj
}

// ConvToStruct 组件转换成实际对象
func (comp Component) ConvToStruct(str interface{}) {
	bytes, _ := json.Marshal(comp.obj)
	json.Unmarshal(bytes, &str)
}
