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
goos: linux
goarch: amd64
pkg: github.com/cbrnrd/ipgen
cpu: AMD Ryzen 5 2600X Six-Core Processor
BenchmarkGenIPv4                44527428                26.54 ns/op            4 B/op          1 allocs/op
BenchmarkGenIPv6                32401519                36.90 ns/op           16 B/op          1 allocs/op
BenchmarkGenIPv4WithExclusions  21224145                56.56 ns/op            4 B/op          1 allocs/op
BenchmarkGenIPv6WithExclusions  16818654                71.12 ns/op           16 B/op          1 allocs/op
```


### Generating 1,000,000 random IPs

Specs:
|Component|Spec|
|---|---|
|CPU|AMD Ryzen 5 2600X Six-Core Processor|
|Arch|x86_64|
|RAM|16GB DDR4|
|OS|Ubuntu WSL|
|Go|1.21.4|

```bash
$ hyperfine --warmup 3 "./ipgen -n 1000000 -o ips.txt" "./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16" "./ipgen -6 -n 1000000 -o ips.txt" "./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16"
Benchmark 1: ./ipgen -n 1000000 -o ips.txt
  Time (mean ± σ):      3.122 s ±  0.016 s    [User: 1.713 s, System: 3.725 s]
  Range (min … max):    3.105 s …  3.152 s    10 runs

Benchmark 2: ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      3.172 s ±  0.032 s    [User: 1.807 s, System: 3.750 s]
  Range (min … max):    3.150 s …  3.255 s    10 runs

Benchmark 3: ./ipgen -6 -n 1000000 -o ips.txt
  Time (mean ± σ):      3.524 s ±  0.103 s    [User: 1.955 s, System: 4.035 s]
  Range (min … max):    3.374 s …  3.758 s    10 runs

Benchmark 4: ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):      3.776 s ±  0.137 s    [User: 2.180 s, System: 4.287 s]
  Range (min … max):    3.519 s …  4.015 s    10 runs

Summary
  ./ipgen -n 1000000 -o ips.txt ran
    1.02 ± 0.01 times faster than ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
    1.13 ± 0.03 times faster than ./ipgen -6 -n 1000000 -o ips.txt
    1.21 ± 0.04 times faster than ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
```