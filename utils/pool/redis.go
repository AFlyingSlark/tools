package pool

import (
	"context"
	"fmt"
	"runtime"

	"github.com/go-redis/redis/v8" // 导入时指定了版本v8,忽略版本是一个常见错误
)

// go-redis 已经实现了 redis 的连接池管理, 因此我们不需要自己手动管理 redis 的连接(获取或放入).
func InitRedisPool(ctx context.Context, addr, password string) *redis.Client {
	fmt.Println("init redis pool")
	cli := redis.NewClient(&redis.Options{
		Addr:         addr,
		DB:           2,
		Password:     password,
		PoolSize:     4 * runtime.NumCPU(), // 连接池最大数量 4*cpu数
		MinIdleConns: 2 * runtime.NumCPU(), // 维持最小闲置连接数
	})

	_, err := cli.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("init redis pool ok")

	return cli
}
