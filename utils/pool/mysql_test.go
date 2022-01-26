package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

func Test_mysqlConnectPool(t *testing.T) {
	db := InitMysqlPool("root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=true")
	var wg sync.WaitGroup
	num := 15 // 大于最大连接数
	wg.Add(num)
	fmt.Println("最大打开连接数测试开始")
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			openMaxConn(db, i)
		}(i)
	}
	wg.Wait()

	fmt.Println("最大空闲连接数测试开始")
	for i := 0; i < 20; i++ {
		idleConn(db)
		time.Sleep(3 * time.Second)
	}
}

func openMaxConn(db *gorm.DB, i int) {
	var connectionID int
	var result int
	err := db.DB().QueryRow("SELECT CONNECTION_ID()").Scan(&connectionID) // 查询连接句柄ID
	if err != nil {
		fmt.Println("query connectionID failed", err.Error())
		return
	}

	fmt.Println("worker:", i, "connectionID:", connectionID)

	err = db.DB().QueryRow("SELECT sleep(10)").Scan(&result) // 连接句柄休眠
	if err != nil {
		fmt.Println("query sleep connectionID failed", err.Error())
		return
	}
}

/*
最大打开连接数测试开始  最大数5个
worker: 14 connectionID: 110376    1
worker: 13 connectionID: 110376
worker: 3 connectionID: 110376
worker: 7 connectionID: 110380     2
worker: 8 connectionID: 110377     3
worker: 0 connectionID: 110378     4
worker: 1 connectionID: 110379     5
worker: 10 connectionID: 110380
worker: 12 connectionID: 110378
worker: 11 connectionID: 110380
worker: 2 connectionID: 110380
worker: 4 connectionID: 110376
worker: 6 connectionID: 110378
*/

func idleConn(db *gorm.DB) {
	var connectionID int

	err := db.DB().QueryRow("SELECT CONNECTION_ID()").Scan(&connectionID) // 查询连接句柄ID
	if err != nil {
		fmt.Println("query connectionID failed", err.Error())
		return
	}

	fmt.Println("connectionID:", connectionID)
}

/*
最大空闲连接数测试开始 最大数3个
connectionID: 110386 1
connectionID: 110387 2
connectionID: 110389 3
connectionID: 110386
connectionID: 110387
connectionID: 110389
connectionID: 110386
connectionID: 110387
connectionID: 110389
connectionID: 110386
connectionID: 110387
connectionID: 110389
connectionID: 110386
connectionID: 110387
connectionID: 110389
connectionID: 110386
connectionID: 110387
connectionID: 110389
*/
