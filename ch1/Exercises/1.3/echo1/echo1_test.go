/*
Cache destination:

	$ go env GOCACHE
	/home/geek/.cache/go-build

Delete all cached data:

	$ go clean -cache


Removes all cached test results (but not cached build results).

    $ go clean -testcache

*/

package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := Echo1()
		b.Log(out)
	}

	/*
		$ go test -bench=.
		goos: linux
		goarch: amd64
		pkg: echo1
		cpu: Intel(R) Core(TM) Ultra 5 225U
			BenchmarkEcho1-14         352692              3779 ns/op
			--- BENCH: BenchmarkEcho1-14
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			echo1_test.go:27: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
			... [output truncated]
			PASS
			ok      echo1   3.302s
	*/
}
