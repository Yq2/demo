package main

import (
	"database/sql"
	"log"
	"time"

	"git.myscrm.cn/golang/common/json"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 驱动注册
	// 所有驱动都是注册在 sql.drivers
	// drivers   = make(map[string]driver.Driver)
	// 所有数据库驱动都要实现sql.Driver接口（该接口只有一个方法Open）
	// _ "github.com/go-sql-driver/mysql" 通过这种方式注册驱动，实际上会在导入驱动包时执行
	//func init() {
	//	sql.Register("postgres", &Driver{})
	//}
	db, err := sql.Open("mysql", "root:123456@tcp(47.112.109.72:3306)/YqStudio?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}

	// 最大总连接数 = 最大打开连接数 + 最大空闲连接数

	// 设置最大的并发打开连接数为5
	// 最大打开连接数并不是总连接数
	// 设置这个数小于或等于0表示没有显示设置，使用默认设置
	db.SetMaxOpenConns(5)

	// 设置最大空闲连接数为5
	// 设置最大空闲连接数大于最大打开连接是没有意义的
	// 最大空闲连接并不是越大越好
	db.SetMaxIdleConns(5)

	// 设置连接可重用的最大时间长度
	// 设置为0意味着没有最大生命周期，连接总是可重用的
	// 这并不能保证连接在连接池中完整的存放一小时，很可能由于某种原因导致连接变得不可用，并在此之前自动关闭
	// 一个连接在创建后任可以使用一个多小时，只是说一个小时后不能再被重用了
	// 这不是空闲超时，连接将在第一次创建后1小时内过期，而不是1小时后变成空闲
	// 每秒自动运行一次清理操作以便从池中删除过期连接
	// 理论上ConnMaxLifeTime越短，从零开始创建连接的频率就越高

	db.SetConnMaxLifetime(time.Hour)

	result, err := db.Query("show tables")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(json.MarshalToStringNoError(result))

}
