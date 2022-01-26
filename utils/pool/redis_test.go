package pool

import (
	"context"
	"fmt"
	"sync"
)

// 连接池测试
func connectPoolTest() {
	ctx := context.TODO()

	fmt.Println("--------connect Pool Test-----------")
	client := InitRedisPool(ctx, "localhost:6379", "")
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				client.Set(ctx, fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(ctx, fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, IdleConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().IdleConns)
		}()
	}

	wg.Wait()
}
