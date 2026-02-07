package defer_test

import (
	"testing"
)

// sum 计算从 0 到 max-1 的整数和
func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}
	return total
}

// openDefer 手动实现的 defer 功能
func openDefer(f func()) {
	f()
}

// WithOpenDefer 使用手动实现的 defer
func WithOpenDefer() {
	openDefer(func() {
		//没有直接调用sum函数,而是使用了匿名函数作为参数
		sum(10)
	})
}

// WithoutDefer 没有使用 defer
func WithoutDefer() {
	sum(10)
}

// BenchmarkWithoutDefer 测试没有使用 defer
func BenchmarkWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithoutDefer()
	}
}

// BenchmarkFooWithDefer 测试原生的 defer
func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {
			sum(10)
		}()
	}
}

// BenchmarkWithOpenDefer 测试开放编码的 defer
func BenchmarkWithOpenDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithOpenDefer()
	}
}

//Unit Tests 单元测试用于测试功能是否正确
//Benchmark Tests 基准测试用于测试代码性能
