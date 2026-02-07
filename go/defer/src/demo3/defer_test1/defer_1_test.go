package defer_test1_test

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

// WithDirect 直接调用
func WithDirect() {

	//模拟一些处理逻辑
	//	//time.Sleep(1 * time.Second)
	sum(10)
	sum(10)
	sum(10)
	sum(10)
	sum(10)
}

// WithOpenDefer 使用开放编码的 defer
func WithOpenDefer() {
	defer sum(10)
	defer sum(10)
	defer sum(10)
	defer sum(10)
	defer sum(10)
	//模拟一些处理逻辑
	//time.Sleep(1 * time.Second)
}

// BenchmarkWithDirect 测试原生的 defer
func BenchmarkWithDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithDirect()
	}
}

// BenchmarkWithOpenDefer 测试开放编码的 defer
func BenchmarkWithOpenDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WithOpenDefer()
	}
}
