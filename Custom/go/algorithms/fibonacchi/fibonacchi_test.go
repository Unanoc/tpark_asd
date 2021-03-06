package fibonacchi

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	if getIter(25) != 75025 {
		fmt.Println(getIter(25))
		t.Error()
	}
}

func TestRecurse(t *testing.T) {
	if getRecurse(25) != 75025 {
		fmt.Println(getRecurse(25))
		t.Error()
	}
}

func benchmarkIter(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		getIter(n)
	}
}

func benchmarkRecurse(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRecurse(n)
	}
}

func BenchmarkIter10(b *testing.B)       { benchmarkIter(10, b) }
func BenchmarkIter100(b *testing.B)      { benchmarkIter(100, b) }
func BenchmarkIter1000(b *testing.B)     { benchmarkIter(1000, b) }
func BenchmarkIter10000(b *testing.B)    { benchmarkIter(10000, b) }
func BenchmarkIter100000(b *testing.B)   { benchmarkIter(100000, b) }
func BenchmarkIter1000000(b *testing.B)  { benchmarkIter(1000000, b) }
func BenchmarkIter10000000(b *testing.B) { benchmarkIter(10000000, b) }

func BenchmarkRecurse10(b *testing.B)     { benchmarkRecurse(10, b) }
func BenchmarkRecurse100(b *testing.B)    { benchmarkRecurse(100, b) }
func BenchmarkRecurse1000(b *testing.B)   { benchmarkRecurse(1000, b) }
func BenchmarkRecurse10000(b *testing.B)  { benchmarkRecurse(10000, b) }
func BenchmarkRecurse100000(b *testing.B) { benchmarkRecurse(100000, b) }
