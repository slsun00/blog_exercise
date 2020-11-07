package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName  = "mysql1"
	Password  = "123"
	Protocol  = "tcp"
	Address   = "127.0.0.1"
	Port      = "3306"
	DBName    = "blog"
	Charset   = "utf8mb48"
	ParseTime = true
	Loc       = "Local"
)

type DB struct{
	db  *gorm.DB
}
// 连接数据库,返回这个数据库操作句柄
func DBOpen() (db  *gorm.DB) {
	// 怎么才能出现检查数据库，发现不存在即自动创建呢？，情况给如何分析
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@%(%s:%s)/%s?chartset=%s&parseTime=%v&loc=%s", UserName,
		Password, Protocol, Address, Port, DBName, Charset, ParseTime, Loc)

	// 新的不用使用 db.Close() ?? ,源码上面也看不到啊
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open fail",err)
		return
	}

	return db
	// 1. 在处理	数据库不存在就初始化一个库
	// 2. 数据库在连接以前就应该确认已经创建了一个数据库
	// 上面两个数据库，但是是该以谁为主呢？
}

// 文章啊的增删改查
// 增删改查的对象是什么？文章？用户？标题？
//func (this *DB)

func AddCategory(name string) error {

}
