```
$ go test -v -bench . -benchmem
=== RUN   TestRawMessageToInt
--- PASS: TestRawMessageToInt (0.00s)
=== RUN   TestRequests
--- PASS: TestRequests (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/erikdubbelboer/json-iterator-benchmark
BenchmarkStandard-16            100000       14680 ns/op      2096 B/op        43 allocs/op
BenchmarkEasyJson-16            200000        6080 ns/op      1808 B/op        38 allocs/op
BenchmarkJsonIterator-16        200000       10843 ns/op      4106 B/op       112 allocs/op
PASS
ok    github.com/erikdubbelboer/json-iterator-benchmark  5.199s
```
