package main

import (
	"testing"
	. "github.com/cbrnrd/ipgen/pkg/ip"
)


func BenchmarkGenIPv4(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenIPv4()
		}
	})
}