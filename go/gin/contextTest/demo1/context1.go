package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务终止:", ctx.Err())
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("过去了1秒")
			}
		}
	}(ctx)
	// 等待协程退出
	time.Sleep(5 * time.Second)
}
