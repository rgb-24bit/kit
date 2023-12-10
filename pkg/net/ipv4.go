package net

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type IPv4Network struct {
	IP      string
	NetWork string
	Mask    string
	IPCount uint32
	FirstIP string
	LastIP  string
}

func ParseIPNetwork(addr string) (*IPv4Network, error) {
	ip, ipNet, err := net.ParseCIDR(addr)
	if err != nil {
		return nil, err
	}

	maskSize, _ := ipNet.Mask.Size()
	network := IPv4StringToInt(ipNet.String())

	ipNetwork := IPv4Network{
		IP:      ip.String(),
		NetWork: ipNet.String(),
		IPCount: IPv4Count(uint32(maskSize)),
		Mask:    IPv4IntToString(0xffffffff << (32 - maskSize)),
		FirstIP: IPv4IntToString(network + 1),
		LastIP:  IPv4IntToString((network | (0xffffffff >> maskSize)) - 1),
	}

	return &ipNetwork, nil
}

func IPv4Count(maskSize uint32) uint32 {
	return 1<<(32-maskSize) - 2 // diff network ip & broadcast ip
}

func IPv4IntToString(val uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", val>>24, (val>>16)&0xff, (val>>8)&0xff, val&0xff)
}

func IPv4StringToInt(val string) uint32 {
	var (
		items        = strings.Split(val, ".")
		ret   uint32 = 0
	)

	for _, item := range items {
		iv, _ := strconv.ParseUint(item, 10, 8)
		ret = ret<<8 + uint32(iv)
	}

	return ret
}
