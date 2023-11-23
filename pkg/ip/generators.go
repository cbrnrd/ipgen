package ip

import (
	"encoding/binary"
	"math/rand"
	"net"
)

// Generate a random IPv4 address.
// The first octet will not be 0.
func GenIPv4() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()

	// check first octet is not 0
	if ip & 0xff000000 == 0 {
		return GenIPv4()
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