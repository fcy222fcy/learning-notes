package main

// 验证defer的类型
func main() {
	//Recycle()
	Condition(true)

	//defer fmt.Println("hello")

	//少量或者多个defer
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
	//defer func() {
	//	var n int
	//	n++
	//}()
}

// Recycle 循环中的defer
//func Recycle() {
//	for i := 0; i < 10; i++ {
//		defer func(i int) {
//		}(i)
//	}
//}

// Condition 条件判断内的defer
func Condition(flag bool) {
	if flag {
		defer func() {

		}()
	}
}
