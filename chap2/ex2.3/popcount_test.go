package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(666)
	}
}

func BenchmarkPopCountRewrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite(666)
	}
}
