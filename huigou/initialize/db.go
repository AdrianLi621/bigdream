package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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
	host := v.Get("Host").(string)
	user := v.Get("UserName").(string)
	password := v.Get("Password").(string)
	port := v.Get("Port").(int)
	database := v.Get("DBName").(string)
	str := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "hg_",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败")
	}
	return db
}
