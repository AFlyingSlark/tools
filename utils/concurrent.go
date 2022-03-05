package utils

import (
	"log"
	"runtime/debug"
	"sync"

	"go.uber.org/zap"
)

/* EnsureGo 并发函数，确保在返回前已经开始执行
参数:
*	functions	...func()
返回值:
*/
func EnsureGo(logger *zap.Logger, functions ...func()) {
	wg := &sync.WaitGroup{}
	wg.Add(len(functions))
	for i := range functions {
		go func(i int) {
			defer func() {
				if err := recover(); err != nil {
					if logger != nil {
						logger.Error("发生panic", zap.Any("错误信息", err), zap.ByteString("堆栈", debug.Stack()))
					} else {
						log.Println("发生panic")
					}
				}
			}()
			wg.Done()

			functions[i]()
		}(i)
	}

	wg.Wait()
}
