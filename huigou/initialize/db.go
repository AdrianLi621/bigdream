package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sync"
)

/**
连接数据库
*/
var DB *gorm.DB
var once sync.Once

/**
初始化数据库
*/
func InitDB() *gorm.DB {
	once.Do(func() {
		if DB == nil {
			DB = ConnectDB()
		}
	})
	return DB
}

/**
加载数据库配置
*/
func LoadDBConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("database")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic("加载database配置文件出错")
	}
	return v
}

/**
连接数据库
*/
func ConnectDB() *gorm.DB {
	v := LoadDBConfig()
	driver := v.Get("DBDriver").(string)
	host := v.Get("Host").(string)
	user := v.Get("UserName").(string)
	password := v.Get("Password").(string)
	port := v.Get("Port").(int)
	database := v.Get("DBName").(string)
	idle := v.Get("MaxIdleConns").(int)
	openconn := v.Get("MaxOpenConns").(int)
	str := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(driver, str)
	if err != nil {
		panic("连接数据库失败")
	}
	db.DB().SetMaxIdleConns(idle)
	db.DB().SetMaxOpenConns(openconn)
	return db
}
