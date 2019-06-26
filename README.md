Benchmarking [json-iterator](https://github.com/json-iterator/go) vs [easyjson](https://github.com/mailru/easyjson)
---

```
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/erikdubbelboer/json-iterator-benchmark
BenchmarkUnmarshalStandard-16                	  100000	     20712 ns/op	    1616 B/op	      43 allocs/op
BenchmarkUnmarshalStandardWithEasy-16        	  100000	     14321 ns/op	    2096 B/op	      43 allocs/op
BenchmarkUnmarshalEasyJson-16                	  200000	      6156 ns/op	    1808 B/op	      38 allocs/op
BenchmarkUnmarshalJsonIterator-16            	  200000	      7643 ns/op	    2721 B/op	     113 allocs/op
BenchmarkUnmarshalJsonIteratorWithEasy-16    	  200000	     11117 ns/op	    4106 B/op	     112 allocs/op
BenchmarkMarshalStandard-16                  	  200000	      7943 ns/op	    2543 B/op	      31 allocs/op
BenchmarkMarshalStandardWithEasy-16          	  200000	     11598 ns/op	    3546 B/op	      21 allocs/op
BenchmarkMarshalEasyJson-16                  	  500000	      3219 ns/op	    1916 B/op	      14 allocs/op
BenchmarkMarshalJsonIterator-16              	  200000	      7256 ns/op	    2227 B/op	      25 allocs/op
BenchmarkMarshalJsonIteratorWithEasy-16      	  100000	     12174 ns/op	    4706 B/op	      22 allocs/op
PASS
ok  	github.com/erikdubbelboer/json-iterator-benchmark	17.742s
```

Seems like for the case of encoding or decoding OpenRTB easyjson is much faster.

To be honest I don't have much trust in json-iterator after seeing [code like this](https://github.com/json-iterator/go/blob/27518f6661eba504be5a7a9a9f6d9460d892ade3/config.go#L323-L358) that casts a `string` to a `[]byte` even though it could probably get away with using `reflect` here to not cause a heap allocation.
