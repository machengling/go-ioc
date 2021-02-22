package test

import (
	"fmt"
	"ioc-container/pkg/store"
	"ioc-container/pkg/store/orm/gorm"
	"testing"
)

func init() {
}

// func TestIocContainer(t *testing.T) {

// 	// 新建一个组件
// 	component := ioc.NewComponent(&common.TestConponent{
// 		Name: "asda",
// 	})

// 	// 注册一个组件
// 	ioc.RegisteComponent(component)
// 	gorm.RegiteMysqlDB(store.DBConfig{
// 		Username: "",
// 		Password: "",
// 		URL:      "",
// 		DBname:   "",
// 	})
// 	// 获取组件
// 	component2 := ioc.NewComponent(&common.TestConponent{})
// 	component2.SetName("mysqldb")
// 	componentReps, _ := ioc.GetComponent(component2)

// 	// 将组件转换成实际结构体
// 	rsp := common.TestConponent{}
// 	componentReps.ConvToStruct(&rsp)
// 	t.Log(rsp, rsp.ConponentName())
// }

func TestIocContainer(t *testing.T) {

	gorm.RegiteMysqlGormDB(store.DBConfig{
		Username: "root",
		Password: "root",
		URL:      "10.10.30.60:3306",
		DBname:   "microservice",
		Options: map[string]string{
			"charset":   "utf8mb4",
			"parseTime": "True",
			"loc":       "Local",
		},
	})
	db := gorm.GetMysqlGormDB()
	var resp interface{}
	db.Table("t_device_4").Find(&resp)
	fmt.Println("resp", resp)
}
