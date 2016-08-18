package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(666)
	}
}

func BenchmarkPopCountRewrite2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRewrite2(666)
	}

}
