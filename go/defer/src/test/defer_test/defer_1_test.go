package defer_test

import "testing"

//使用defer和没有defer

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}
	return total
}
func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}
func fooWithoutDefer() {
	sum(10)
}

// 有defer
func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

// 没有defer
func BenchmarkFooWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}
}
