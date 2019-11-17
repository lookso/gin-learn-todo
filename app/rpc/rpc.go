/*
@Time : 2019-11-17 15:37 
@Author : Tenlu
@File : rpc
@Software: GoLand
*/
package rpc

import (
	"fmt"
	"net"
)

type Service struct {
	Port string
}

func (r *Service) StartRpc() *Service {
	addr := fmt.Sprintf("%s:%s", GetLocalIP(), r.Port)

	fmt.Printf("addr:", addr+":"+r.Port)

	return &Service{}
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
