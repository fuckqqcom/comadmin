package config

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
	"log"
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
		Db       int
	}
	Es struct {
		Host  string //es host
		Index string // es index
	}
}

var (
	EngDb       *xorm.Engine
	RedisClient *redis.Client
	EsClient    *elastic.Client
	EsIndex     string
)

func NewConfig(path string) (config Config) {
	Load(path, &config)
	config.loadDb()
	config.LoadRedis()
	config.LoadElastic()
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
func (c *Config) LoadRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Dns,
		Password:     "",
		DB:           c.Redis.Db,
		MaxRetries:   1,
		PoolSize:     c.Redis.PoolSize,
		MinIdleConns: c.Redis.MinIdle,
	})
	result, err := RedisClient.Ping().Result()
	log.Printf("redis conn %v %v", result, err)
}

func (c *Config) LoadElastic() {
	var err error
	EsIndex = c.Es.Index
	EsClient, err = elastic.NewSimpleClient(elastic.SetURL(c.Es.Host))
	if err != nil {
		log.Printf("elastic conn is error %s", err)
	}

	//info := elasticsearch.Config{
	//	Addresses: []string{""},
	//	Transport: &http.Transport{
	//		MaxIdleConnsPerHost:   10,
	//		ResponseHeaderTimeout: 500 * time.Millisecond,
	//		DialContext:           (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
	//		TLSClientConfig: &tls.Config{
	//			MinVersion: tls.VersionTLS11,
	//			// ...
	//		},
	//	},
	//}
	//if EsClient, err = elasticsearch.NewClient(info); err != nil {
	//	log.Printf("elastic conn is error %v", err)
	//}
}
