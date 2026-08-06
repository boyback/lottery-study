package datasource

import (
	"fmt"
	"log"
	"lottery-study/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil {
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil {
		return masterInstance
	}
	return NewDbMaster()
}

func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", conf.DbMaster.Username, conf.DbMaster.Password, conf.DbMaster.Host, conf.DbMaster.Port, conf.DbMaster.Database, conf.DbMaster.Charset)
	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbHelper NewEngine err:", err)
		return nil
	}
	instance.ShowSQL(true)
	masterInstance = instance
	return masterInstance
}
