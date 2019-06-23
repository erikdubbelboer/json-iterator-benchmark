Benchmarking [json-iterator](https://github.com/json-iterator/go) vs [easyjson](https://github.com/mailru/easyjson)
---

```
$ go test -v -bench . -benchmem
=== RUN   TestRawMessageToInt
--- PASS: TestRawMessageToInt (0.00s)
=== RUN   TestUnmarshal
--- PASS: TestUnmarshal (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/erikdubbelboer/json-iterator-benchmark
BenchmarkUnmarshalStandard-16        100000	   14652 ns/op	 2096 B/op	   43 allocs/op
BenchmarkUnmarshalEasyJson-16        200000	    6140 ns/op	 1808 B/op	   38 allocs/op
BenchmarkUnmarshalJsonIterator-16    200000	   11588 ns/op	 4106 B/op	  112 allocs/op
BenchmarkMarshalStandard-16          100000	   12305 ns/op	 3545 B/op	   21 allocs/op
BenchmarkMarshalEasyJson-16          500000	    3425 ns/op	 1916 B/op	   14 allocs/op
BenchmarkMarshalJsonIterator-16      100000	   12578 ns/op	 4703 B/op	   22 allocs/op
PASS
ok  	github.com/erikdubbelboer/json-iterator-benchmark	9.849s
```

Seems like for the case of encoding or decoding OpenRTB easyjson is much faster.

To be honest I don't have much trust in json-iterator after seeing [code like this](https://github.com/json-iterator/go/blob/27518f6661eba504be5a7a9a9f6d9460d892ade3/config.go#L323-L358) that casts a `string` to a `[]byte` even though it could probably get away with using `reflect` here to not cause a heap allocation.
