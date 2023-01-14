package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	ip, err := GetIntranetIP("eno1")
	if nil != err {
		fmt.Println(1111)
	}
	fmt.Println(ip)

}

// 获取本机内网IP
func GetIntranetIP(name string) (ip string, err error) {
	// 获取内网IP信息
	interfaces, err := net.InterfaceByName(name)
	fmt.Println(interfaces)
	if err != nil {
		return
	}
	addrs, err := interfaces.Addrs()
	fmt.Println(addrs)
	if err != nil {
		return
	}
	for _, addr := range addrs {
		fmt.Println(addr.String())
		ipInfo := strings.Split(addr.String(), "/")
		fmt.Println(ipInfo)
		if len(strings.Split(ipInfo[0], ".")) == 4 {
			ip = ipInfo[0]
			fmt.Println(ip)
			return
		}
	}
	return
}
