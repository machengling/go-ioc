package gorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/machengling/go-ioc/pkg/ioc"
	"github.com/machengling/go-ioc/pkg/store"

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

/*
三色标记法（默认所有的对象都是白色）
（1）stop the world
（2）找到所有的roots，标记为灰色，并加入队列
（3）开启协程 gc write barrier（为后面监测被修改对象时，重新加入队列做处理）

（4）start the world
（5）将队列中的对象取出，并标记为黑色。
（6）同时如果对象持有其他对象的指针引用，则将持有的对象加入到队列中
（7）期间如果用户代码修改对象，那么会触发写屏障，将对象标记为灰色，并加入单独的扫描队列中

（8）stop the world
（9）将触发写屏障后，加入的列表，遍历出来，重新标记
（10）start the world

（11）清理标记为白色的对象
*/
