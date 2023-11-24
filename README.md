# `ipgen`

Generate lots of IPs, pretty quickly.

## Usage

```bash
$ ipgen -h
Usage: ./ipgen [options]
Generates random IPv4 addresses and writes them to a stream
Version: 0.1.0
  -6	Generate IPv6 addresses instead of IPv4 addresses (default: false)
  -c int
    	Number of workers to use to generate the data (default 4)
  -n int
    	Number of IPs to generate. -1 means infinite (default -1)
  -o string
    	Output path for the generated data (default "-")
  -x string
    	IP ranges to exclude from the generated data (format: range,range,range)
```

## Benchmarks

### Generation function benchmark

```bash
$ GOMAXPROCS=1 go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/cbrnrd/ipgen
BenchmarkGenIPv4               	28962502	        41.33 ns/op	      15 B/op	       1 allocs/op
BenchmarkGenIPv6               	 6986734	       170.1 ns/op	      63 B/op	       2 allocs/op
BenchmarkGenIPv4WithExclusions 	 9832515	       120.8 ns/op	      15 B/op	       1 allocs/op
BenchmarkGenIPv6WithExclusions 	 2491005	       480.8 ns/op	      64 B/op	       2 allocs/op
```


### Generating 1,000,000 random IPs

```bash
$ hyperfine --warmup 3 "./ipgen -n 1000000 -o ips.txt" "./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16" "./ipgen -6 -n 1000000 -o ips.txt" "./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16"
Benchmark 1: ./ipgen -n 1000000 -o ips.txt
  Time (mean ± σ):      2.274 s ±  0.071 s    [User: 1.686 s, System: 2.391 s]
  Range (min … max):    2.200 s …  2.432 s    10 runs

Benchmark 2: ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      2.372 s ±  0.047 s    [User: 1.782 s, System: 2.470 s]
  Range (min … max):    2.311 s …  2.459 s    10 runs

Benchmark 3: ./ipgen -6 -n 1000000 -o ips.txt
  Time (mean ± σ):      2.541 s ±  0.061 s    [User: 1.871 s, System: 2.579 s]
  Range (min … max):    2.466 s …  2.663 s    10 runs

Benchmark 4: ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      2.799 s ±  0.053 s    [User: 2.400 s, System: 2.705 s]
  Range (min … max):    2.724 s …  2.876 s    10 runs

Summary
  ./ipgen -n 1000000 -o ips.txt ran
    1.04 ± 0.04 times faster than ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
    1.12 ± 0.04 times faster than ./ipgen -6 -n 1000000 -o ips.txt
    1.23 ± 0.04 times faster than ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
```