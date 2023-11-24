package ip

import (
	"encoding/binary"
	"math/rand"
	"net"
)

// Generate a random IPv4 address.
// The IP is guaranteed to be a valid, non loopback address
// or private address.
func GenIPv4() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()
	o1, o2 := byte(ip>>24)&0xff, byte(ip>>16)&0xff

	for (o1 == 0) || // 0.0.0.0/8 - Invalid address
		(o1 == 127) || // 127.0.0.0/8 - Loopback
		(o1 >= 224) || // 224.*.*.*+ - Multicast
		(o1 == 10) || // 10.0.0.0/8 - Internal network
		(o1 == 192 && o2 == 168) || // 192.168.0.0/16 - Internal network
		(o1 == 172 && o2 >= 16 && o2 < 32) { // 172.16.0.0/14 - Internal network
		ip = rand.Uint32()
		o1, o2 = byte(ip>>24)&0xff, byte(ip>>16)&0xff
	}

	binary.LittleEndian.PutUint32(buf, ip)
	return net.IP(buf).String()
}

// Generate a random IPv6 address.
func GenIPv6() string {
	size := 16
	ip := make([]byte, size)
	for i := 0; i < size; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return net.IP(ip).To16().String()
}

// Checks if an IP is in any of the excluded ranges
func IsExcluded(ip string, excludedRanges []net.IPNet) bool {
	for _, network := range excludedRanges {
		if network.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}

func GenIPv4WithExclusions(excludedRanges []net.IPNet) string {
	ip := GenIPv4()
	if IsExcluded(ip, excludedRanges) {
		return GenIPv4WithExclusions(excludedRanges)
	}
	return ip
}

func GenIPv6WithExclusions(excludedRanges []net.IPNet) string {
	ip := GenIPv6()
	if IsExcluded(ip, excludedRanges) {
		return GenIPv6WithExclusions(excludedRanges)
	}
	return ip
}
