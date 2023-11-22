package ip

import (
	"encoding/binary"
	"math/rand"
	"net"
)

func GenIPv4() string {
	buf := make([]byte, 4)
	ip := rand.Uint32()

	binary.LittleEndian.PutUint32(buf, ip)
	return net.IP(buf).String()
}