package pool

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// gorm底层是根据sql.DB实现的，而sql.DB里面提供了相关的配置函数.执行sql语句才会使用连接池句柄

// InitMysqlPool 初始化MySQL连接池
func InitMysqlPool(slog *log.Logger, dataSource string) *gorm.DB {
	var err error
	var newLogger logger.Interface

	if slog == nil {
		// 无日志器则使用gorm默认日志器,info打印级别
		newLogger = logger.Default.LogMode(logger.Info)
	} else {
		// 使用自定义日志器
		newLogger = logger.New(
			slog, // 实现io.Writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		)
	}

	DB, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置表相关的名称为单数
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		defer sqlDB.Close()
		panic(err)
	}

	sqlDB.SetMaxOpenConns(5) // 设置最大连接数. 默认0无限制  如果不根据服务端的最大连接数设置的话，并发起来的时候可能会报Too many connections,值不是设置的越大越好的。根据服务器的配置进行相应的设置才能使得吞吐量最大化。
	sqlDB.SetMaxIdleConns(3) // 设置最大空闲连接数

	return DB
}
