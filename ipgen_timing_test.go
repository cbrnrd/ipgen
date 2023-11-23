package main

import (
	"net"
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

func BenchmarkGenIPv6(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenIPv6()
		}
	})
}

func BenchmarkGenIPv4WithExclusions(b *testing.B) {
	excludedRanges := []net.IPNet{
		{
			IP:   net.ParseIP("192.168.0.0"),
			Mask: net.CIDRMask(16, 32),
		},
		{
			IP:   net.ParseIP("10.0.0.0"),
			Mask: net.CIDRMask(16, 32),
		},
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenIPv4WithExclusions(excludedRanges)
		}
	})
}

func BenchmarkGenIPv6WithExclusions(b *testing.B) {
	excludedRanges := []net.IPNet{
		{
			IP:   net.ParseIP("192.168.0.0"),
			Mask: net.CIDRMask(16, 32),
		},
		{
			IP:   net.ParseIP("10.0.0.0"),
			Mask: net.CIDRMask(16, 32),
		},
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GenIPv6WithExclusions(excludedRanges)
		}
	})
}
