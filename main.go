package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {

	oneMb := 1024 * 1024
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

	flag.IntVar(&n, "n", -1, "Number of IPs to generate. -1 means infinite")
	flag.StringVar(&outpath, "o", "-", "Output path for the generated data")
	flag.StringVar(&excludedRangesStr, "x", "", "IP ranges to exclude from the generated data (format: range,range,range)")
	flag.BoolVar(&v6, "6", false, "Generate IPv6 addresses instead of IPv4 addresses (default: false)")

	flag.Parse()

	excludedRanges := ParseExcludedRanges(excludedRangesStr)
	outFile := SetupOutput(outpath)
	ipBytesMultiplier := 1
	if v6 {
		ipBytesMultiplier = 4
	}

	buf := make([]byte, 0, oneMb)
	bufLen := 0
	generateRandomIP := func() net.IP {
		ip := make(net.IP, 4*ipBytesMultiplier)
		rand.Read(ip)
		return ip
	}

	for i := 0; n == -1 || i < n; i++ {
		ip := generateRandomIP()
		if len(excludedRanges) > 0 && IsExcluded(ip, excludedRanges) {
			i--
			continue
		}

		ipStr := ip.String()
		ipBytes := append([]byte(ipStr), '\n')
		if bufLen+len(ipStr) > oneMb {
			outFile.Write(buf)
			
			// reset buffer
			buf = buf[:0]
			bufLen = 0
		}

		buf = append(buf, ipBytes...)
		bufLen += len(ipBytes)
	}

	outFile.Write(buf)

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

// Checks if an IP is in any of the excluded ranges
func IsExcluded(ip net.IP, excludedRanges []net.IPNet) bool {
	for _, network := range excludedRanges {
		if network.Contains(ip) {
			return true
		}
	}
	return false
}
