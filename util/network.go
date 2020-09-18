package util

import (
	"errors"
	"net"
)

func ExternalIP() (net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, _interface := range interfaces {
		if _interface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if _interface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		address, err := _interface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range address {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
