package main

import "fmt"

// 延迟函数操作主函数的返回值
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())

}
func f1() (result int) {
	i := 1
	defer func() {
		result++
	}()
	return i
}

func f2() (result int) {
	i := 1
	defer func() {
		i++
	}()
	return i
}

func f3() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
