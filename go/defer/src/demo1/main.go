package main

import "fmt"

func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	_, _ = f5()
	f6()
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func f5() (int, error) {
	defer fmt.Println("defer")
	return fmt.Println("return")
}

func f6() {
	num := 0
	defer func(n int) {
		fmt.Println("defer:", n)
	}(num)
	//等同于defer fmt.Println("defer:",num)
	for i := 0; i < 10; i++ {
		num++
	}
	fmt.Println(num)
}

//10
//defer: 0

//fmt.Println函数的签名是:
//func Println(a ...interface{})(n int,err error) n代表打印的字节数,err代表可能发生的错误

// 这里可以作为堆的演示
func DeferDemo1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func demo() {
	// 其他业务逻辑
	// 编译器自动插入defer函数，顺序与声明相反
	fmt.Println("C") // 后声明的先执行
	fmt.Println("B")
	fmt.Println("A")
	return // 直接返回，无链表操作
}
