package pool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// gorm底层是根据sql.DB实现的，而sql.DB里面提供了相关的配置函数.执行sql语句才会使用连接池句柄

// InitMysqlPool 初始化MySQL连接池
func InitMysqlPool(dataSource string) *gorm.DB {
	var err error
	DB, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err := DB.DB().Ping(); err != nil {
		defer DB.Close()
		panic(err)
	}

	DB.DB().SetMaxOpenConns(5) // 设置最大连接数. 默认0无限制  如果不根据服务端的最大连接数设置的话，并发起来的时候可能会报Too many connections,值不是设置的越大越好的。根据服务器的配置进行相应的设置才能使得吞吐量最大化。
	DB.DB().SetMaxIdleConns(3) // 设置最大空闲连接数
	DB.SingularTable(true)
	DB.LogMode(true)

	return DB
}
