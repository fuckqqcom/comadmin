package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

type Config struct {
	Db struct {
		Host    string
		User    string
		Pwd     string
		Db      string
		Show    bool
		Port    int
		MaxOpen int
		MaxIdle int
	}
	Redis struct {
		Dns      string
		MinIdle  int
		PoolSize int
	}
}

var (
	EngDb *xorm.Engine
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	return
}

func NewDb() *xorm.Engine {
	return EngDb
}

func (c *Config) loadDb() {
	fmt.Println(c.Db)
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Db.User,
		c.Db.Pwd,
		c.Db.Host,
		c.Db.Db)

	EngDb, err = xorm.NewEngine("mysql", dns)

	ping := EngDb.Ping()
	if ping != nil || err != nil {
		panic(ping)
	}
	EngDb.SetMaxIdleConns(c.Db.MaxIdle)
	EngDb.SetMaxOpenConns(c.Db.MaxOpen)
	EngDb.ShowSQL(c.Db.Show)

}
