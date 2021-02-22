package utils

import (
	"fmt"
	"reflect"
)

// GetComponentName 获取组件的自定义名称
func GetComponentName(obj interface{}) string {
	objtype := reflect.TypeOf(obj)
	objval := reflect.ValueOf(obj)

	fmt.Println("Kind", objtype.Kind())
	if objtype.Kind().String() != reflect.Ptr.String() {
		panic("[ioc-container] param is not ptr")
	}
	strName := objtype.Name()

	methodNum := objtype.NumMethod()

	for i := 0; i < methodNum; i++ {
		methodName := objtype.Method(i).Name
		if methodName == "ConponentName" {
			vals := objval.Method(i).Call(nil)

			if len(vals) == 0 {
				panic("[ioc-container] func ConponentName return nil")
			}
			fmt.Println("vals", vals[0].String())
			strName = vals[0].String()
			break
		}
	}

	return strName
}
