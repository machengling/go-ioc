package gorm

import (
	"fmt"
	"ioc-container/pkg/ioc"
	"ioc-container/pkg/store"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// RegiteMysqlGormDB ..
func RegiteMysqlGormDB(dbconfig store.DBConfig) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	dbpath := getPath(dbconfig)
	db, err := gorm.Open(mysql.Open(dbpath), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("connect to db error: " + err.Error())
	}
	component := ioc.NewComponent(db)
	component.SetName("mysqldb")
	ioc.RegisteComponent(component)

	return
}

// GetMysqlGormDB 获取db实例
func GetMysqlGormDB() *gorm.DB {
	component2 := ioc.NewComponent(&gorm.DB{})
	component2.SetName("mysqldb")
	componentReps, _ := ioc.GetComponent(component2)

	// 将组件转换成实际结构体
	obj := componentReps.GetObj()
	db := obj.(*gorm.DB)
	return db
}

func getPath(dbconfig store.DBConfig) string {
	optionStr := ""
	for k, v := range dbconfig.Options {
		optionStr += (k + "=" + v + "&")
	}
	path := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", dbconfig.Username, dbconfig.Password, dbconfig.URL, dbconfig.DBname, optionStr)
	fmt.Println("path", path)
	return path
}
