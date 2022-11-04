package models

import (
	"cloud-disk/core/internal/config"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	xormLogFile, err := os.OpenFile("logs/xorm_sql.log", os.O_APPEND|os.O_WRONLY, 6)
	if err != nil {
		logx.Errorf("Open xorm_sql.log failed:%v", err)
		return nil
	}
	engine.SetLogger(xlog.NewSimpleLogger(xormLogFile)) // 将日志重定向到文件中
	engine.ShowSQL(true)                                // 打印出生成的SQL语句
	engine.Logger().SetLevel(xlog.LOG_DEBUG)            // 打印调试及以上的信息
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
