package database

import (
	"dl/models"
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"xorm.io/core"
)

/**
 * 实例化数据库引擎
 */
func NewEngine() *xorm.Engine {
	/**
	initConfig := config.InitConfig()

	if initConfig == nil {
		return nil
	}

	engine, err := xorm.NewEngine(initConfig.DBDriver, initConfig.DBSource)
	*/

	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres", "123456@abc", "127.0.0.1", "5432", "dl")
	// dbConf.User, passwd, host, port, database
	engine, err := xorm.NewEngine("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	engine.SetMapper(core.GonicMapper{})

	err = engine.Sync2(
		new(models.User),
		new(models.Course),
		new(models.Chapter),
		new(models.Progress),
		new(models.Chat),
		new(models.Question),
		new(models.Choice),
		new(models.Exercise),
	)
	if err != nil {
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(100)

	return engine
}
