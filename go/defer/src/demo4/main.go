package main

import (
	"fmt"
	"time"
)

// defer预计算
func main() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))
	time.Sleep(time.Second)
}

// 0s

//func main() {
//	startedAt := time.Now()
//	defer func() {
//		fmt.Println(time.Since(startedAt))
//	}()
//	time.Sleep(time.Second)
//}

//1.0005916s
