/*
Cache destination:

	$ go env GOCACHE
	/home/geek/.cache/go-build

Delete all cached data:

	$ go clean -cache


Remove all cached test results (but not cached build results).

    $ go clean -testcache

*/

package main

import "testing"

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out := Echo2()
		b.Log(out)
	}
	/*
		$ go test -bench=.
		goos: linux
		goarch: amd64
		pkg: echo2
		cpu: Intel(R) Core(TM) Ultra 5 225U
		BenchmarkEcho1-14         353331              3381 ns/op
		--- BENCH: BenchmarkEcho1-14
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		    echo2_test.go:25: -test.paniconexit0 -test.timeout=10m0s -test.bench=.
		        ... [output truncated]
		PASS
		ok      echo2   3.135s
	*/
}
