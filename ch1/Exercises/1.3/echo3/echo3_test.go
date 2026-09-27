package main

import (
	"testing"
)

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := Echo3()
		b.Log(out)
	}
	/*
		$ go test -bench=.
		goos: linux
		goarch: amd64
		pkg: echo3
		cpu: Intel(R) Core(TM) Ultra 5 225U
			BenchmarkEcho3-14         365608              3277 ns/op
			--- BENCH: BenchmarkEcho3-14
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo3_test.go:10: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			... [output truncated]
			PASS
			ok      echo3   3.127s
			geek@ThinkPad-T14:~/GoPro/gopl/ch1/Exercises/1.3/echo3$

	*/
}
