# `ipgen`

Generate lots of IPs, pretty quickly.

## Usage

```bash
$ ipgen -h
Usage: ./ipgen [options]
Generates random IPv4 addresses and writes them to a stream
Version: 0.1.0
  -6	Generate IPv6 addresses instead of IPv4 addresses (default: false)
  -n int
    	Number of IPs to generate. -1 means infinite (default -1)
  -o string
    	Output path for the generated data (default "-")
  -x string
    	IP ranges to exclude from the generated data (format: range,range,range)
```

## Benchmarks

### Throughput

```bash
$ ./ipgen | pv > /dev/null
 378MiB 0:00:09 [41.5MiB/s] [     <=>      ]
```

### Generating 1,000,000 random IPs

Specs:
|Component|Spec|
|---|---|
|CPU|Apple M1 Max|
|Arch|arm64|
|RAM|64GB DDR4|
|OS|macOS Monterey 12.6.5|
|Go|1.21.4|

```bash
$ hyperfine --warmup 3 "./ipgen -n 1000000 -o ips.txt" "./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16" "./ipgen -6 -n 1000000 -o ips.txt" "./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16"
Benchmark 1: ./ipgen -n 1000000 -o ips.txt
  Time (mean ± σ):     326.4 ms ±   5.7 ms    [User: 199.8 ms, System: 127.4 ms]
  Range (min … max):   321.6 ms … 341.4 ms    10 runs
 
Benchmark 2: ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):     332.8 ms ±   1.6 ms    [User: 207.5 ms, System: 126.3 ms]
  Range (min … max):   330.9 ms … 336.2 ms    10 runs
 
Benchmark 3: ./ipgen -6 -n 1000000 -o ips.txt
  Time (mean ± σ):     392.7 ms ±   1.5 ms    [User: 255.8 ms, System: 140.3 ms]
  Range (min … max):   390.3 ms … 394.6 ms    10 runs
 
Benchmark 4: ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
  Time (mean ± σ):     401.7 ms ±   2.3 ms    [User: 265.2 ms, System: 140.3 ms]
  Range (min … max):   399.4 ms … 406.4 ms    10 runs
 
Summary
  ./ipgen -n 1000000 -o ips.txt ran
    1.02 ± 0.02 times faster than ./ipgen -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
    1.20 ± 0.02 times faster than ./ipgen -6 -n 1000000 -o ips.txt
    1.23 ± 0.02 times faster than ./ipgen -6 -n 1000000 -o ips.txt -x 192.168.0.0/16,10.0.0.0/16
```