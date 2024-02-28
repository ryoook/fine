package util

import (
	_ "github.com/joho/godotenv/autoload"
	"net"
)

const (
	defaultHostIP = "0.0.0.0"
)

var hostIP string

func init() {
	hostIP = getHostIP()
}

// HostIP .
func HostIP() string {
	return hostIP
}

func getHostIP() string {
	hostIP = defaultHostIP
	addresses, _ := net.InterfaceAddrs()
	for _, address := range addresses {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				hostIP = ipNet.IP.String()
				break
			}
		}
	}
	return hostIP
}
