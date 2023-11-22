package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	. "github.com/cbrnrd/ipgen/pkg/ip"
)

type IPRangeList struct {
	value []string
}

func (f *IPRangeList) String() string {
	return ""
}

func (f *IPRangeList) Set(s string) error {
	f.value = append(f.value, s)
	return nil
}


func main() {

	concurrency := 20
	n := -1
	outpath := "-"
	excludedRanges := IPRangeList{}

	flag.IntVar(&concurrency, "c", 20, "Number of workers to use to generate the data")
	flag.IntVar(&n, "n", -1, "Number of IPs to generate. -1 means infinite")
	flag.StringVar(&outpath, "o", "-", "Output path for the generated data")
	flag.Var((*IPRangeList)(&excludedRanges), "x", "IP ranges to exclude from the generated data")

	flag.Parse()

	fmt.Println("Excluded ranges:", excludedRanges)

	outFile := os.Stdout
	if outpath != "-" {
		var err error
		outFile, err = os.Create(outpath)
		if err != nil {
			panic(err)
		}
	}
	jobs := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go RunIPv4Routine(&wg, jobs, outFile)
	}

	for i := 0; i < n || n == -1; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}

// Function intended to be invoked via `go`, grabs jobs from the input channel and writes
// a new random IP to the output file
func RunIPv4Routine(wg *sync.WaitGroup, jobs <-chan int, outFile *os.File) {
	defer wg.Done()
	for range jobs {
		outFile.WriteString(GenIPv4() + "\n")
	}
}
