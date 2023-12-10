package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"

	. "github.com/cbrnrd/ipgen/pkg/ip"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	concurrency := 20
	n := -1
	outpath := "-"
	excludedRangesStr := ""
	v6 := false

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Generates random IPv4 addresses and writes them to a stream\n")
		fmt.Fprintf(os.Stderr, "Version: %s (%s) at %s\n\n", version, commit, date)
		flag.PrintDefaults()
	}

	flag.IntVar(&concurrency, "c", 4, "Number of workers to use to generate the data")
	flag.IntVar(&n, "n", -1, "Number of IPs to generate. -1 means infinite")
	flag.StringVar(&outpath, "o", "-", "Output path for the generated data")
	flag.StringVar(&excludedRangesStr, "x", "", "IP ranges to exclude from the generated data (format: range,range,range)")
	flag.BoolVar(&v6, "6", false, "Generate IPv6 addresses instead of IPv4 addresses (default: false)")

	flag.Parse()

	excludedRanges := ParseExcludedRanges(excludedRangesStr)
	outFile := SetupOutput(outpath)

	jobs := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go Run(&wg, jobs, outFile, GetGenerator(v6, excludedRanges))
	}

	for i := 0; i < n || n == -1; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()

}

func SetupOutput(outpath string) *os.File {
	outFile := os.Stdout
	if outpath != "-" {
		var err error
		outFile, err = os.Create(outpath)
		if err != nil {
			panic(err)
		}
	}
	return outFile
}

func ParseExcludedRanges(excludedRangesStr string) []net.IPNet {
	var excludedRanges []net.IPNet = make([]net.IPNet, 0)

	if len(excludedRangesStr) > 0 {
		ranges := strings.Split(excludedRangesStr, ",")
		for _, excludedRange := range ranges {
			_, network, err := net.ParseCIDR(excludedRange)
			if err != nil {
				panic(err)
			}
			excludedRanges = append(excludedRanges, *network)
		}
	}
	return excludedRanges
}

// Runs `generatorFunc` and writes the result to `outFile`.
// Intended to be invoked as a goroutine.
func Run(wg *sync.WaitGroup, jobs <-chan int, out io.Writer, generatorFunc func() net.IP) {
	defer wg.Done()
	for range jobs {
		io.WriteString(out, generatorFunc().String()+"\n")
	}
}

// Runs `generatorFunc` and writes the result to `outFile`, excluding any IPs in `excludedRanges`.
// Intended to be invoked as a goroutine.
func RunWithExclusions(wg *sync.WaitGroup, jobs <-chan int, out io.Writer, excludedRanges []net.IPNet, generatorFunc func() net.IP) {
	defer wg.Done()
	for range jobs {
		ip := generatorFunc()
		if !IsExcluded(ip, excludedRanges) {
			io.WriteString(out, ip.String()+"\n")

		}
	}
}

// Returns the correct generator function based on the v6 flag
func GetGenerator(v6 bool, excludedRanges []net.IPNet) func() net.IP {
	if v6 {
		if len(excludedRanges) > 0 {
			return func() net.IP {
				return GenIPv6WithExclusions(excludedRanges)
			}
		}
		return GenIPv6
	}
	if len(excludedRanges) > 0 {
		return func() net.IP {
			return GenIPv4WithExclusions(excludedRanges)
		}
	}
	return GenIPv4
}
