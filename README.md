# `ipgen`

Generate lots of IPs, really fast.

## Benchmarks

### Generation function benchmark

```bash
$ GOMAXPROCS=1 go test -bench=. github.com/cbrnrd/ipgen
goos: darwin
goarch: arm64
pkg: github.com/cbrnrd/ipgen
BenchmarkGenIPv4 	30954295	        39.36 ns/op
PASS
ok  	github.com/cbrnrd/ipgen	3.248s
```


### Generating 1,000,000 random IPs:

```bash
$ hyperfine --warmup=3 "./ipgen -n 1000000 -o ips.txt"
Benchmark 1: ./ipgen -n 1000000 -o ips.txt
  Time (mean ± σ):      2.189 s ±  0.046 s    [User: 1.642 s, System: 2.320 s]
  Range (min … max):    2.138 s …  2.277 s    10 runs
```