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
BenchmarkGenIPv4               	27612655	        40.74 ns/op	      15 B/op	       1 allocs/op
BenchmarkGenIPv6               	 6622725	       178.2 ns/op	      63 B/op	       2 allocs/op
BenchmarkGenIPv4WithExclusions 	 9639518	       123.8 ns/op	      15 B/op	       1 allocs/op
BenchmarkGenIPv6WithExclusions 	 2410920	       500.9 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/cbrnrd/ipgen	5.814s
```


### Generating 1,000,000 random IPs

```bash
$ hyperfine --warmup 3 "./ipgen -n 1000000 -o ips.txt" "./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16" "./ipgen -6 -n 1000000 -o ips.txt" "./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16"
Benchmark 1: ./ipgen -n 1000000 -o ips.txt
  Time (mean ± σ):      2.276 s ±  0.067 s    [User: 1.669 s, System: 2.394 s]
  Range (min … max):    2.206 s …  2.391 s    10 runs

Benchmark 2: ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      2.409 s ±  0.068 s    [User: 1.797 s, System: 2.486 s]
  Range (min … max):    2.336 s …  2.578 s    10 runs

Benchmark 3: ./ipgen -6 -n 1000000 -o ips.txt
  Time (mean ± σ):      2.549 s ±  0.049 s    [User: 1.842 s, System: 2.590 s]
  Range (min … max):    2.487 s …  2.670 s    10 runs

Benchmark 4: ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      2.861 s ±  0.082 s    [User: 2.374 s, System: 2.759 s]
  Range (min … max):    2.781 s …  3.064 s    10 runs

Summary
  ./ipgen -n 1000000 -o ips.txt ran
    1.06 ± 0.04 times faster than ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
    1.12 ± 0.04 times faster than ./ipgen -6 -n 1000000 -o ips.txt
    1.26 ± 0.05 times faster than ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
```