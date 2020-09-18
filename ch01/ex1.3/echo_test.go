// ex1.3 benchmarks various methods for printing command line arguments
package echo_test

import (
	echo "gopl/ch01/ex1.3"
	"testing"
)

var osArgs = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func BenchmarkEchoLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Loop(osArgs)
	}
}

func BenchmarkEchoRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Range(osArgs)
	}
}

func BenchmarkEchoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo.Join(osArgs)
	}
}
