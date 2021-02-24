package utils

import (
	"fmt"
	"reflect"
)

// GetComponentName 获取组件的自定义名称
func GetComponentName(obj interface{}) string {
	if obj == nil {
		panic("[github.com/machengling/go-ioc] GetComponentName param is nil")
	}
	objtype := reflect.TypeOf(obj)
	objval := reflect.ValueOf(obj)

	fmt.Println("Kind", objtype.Kind())
	if objtype.Kind().String() != reflect.Ptr.String() {
		panic("[github.com/machengling/go-ioc] param is not ptr")
	}
	strName := objtype.Name()

	methodNum := objtype.NumMethod()

	for i := 0; i < methodNum; i++ {
		methodName := objtype.Method(i).Name
		if methodName == "ConponentName" {
			vals := objval.Method(i).Call(nil)

			if len(vals) == 0 {
				panic("[github.com/machengling/go-ioc] func ConponentName return nil")
			}
			fmt.Println("vals", vals[0].String())
			strName = vals[0].String()
			break
		}
	}

	return strName
}
