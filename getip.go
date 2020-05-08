package workertools

import (
	"net"
	"strings"

	"fmt"
)

//host ip
//func Get_hostip() string {
//	addrs, err := net.InterfaceAddrs()
//	var ip string
//	if err == nil {
//		for _, a := range addrs {
//			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//				if ipnet.IP.To4() != nil {
//					//log.Infof("get_hostip:%s", ipnet.IP.String())
//					ip += ipnet.IP.String()
//					//return ipnet.IP.String()
//				}
//			}
//		}
//	}
//	return ip
//	//remoteid ,_ := uuid.NewRandom()
//	//return remoteid.String()
//}
func Get_hostip() string {
	addrs, err := net.InterfaceAddrs()
	var ip []string
	var ipret string
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					//log.Infof("get_hostip:%s", ipnet.IP.String())
					ip = append(ip,ipnet.IP.String())
					//return ipnet.IP.String()
				}
			}
		}
	}
	var ipnil string
	for _,ipmem := range ip{//优先选择10网段
		dataindex := strings.Index(ipmem,"10")
		if dataindex==0{//10.xxx
			//fmt.Println(dataindex)
			ipret = ipmem
		} else {
			ipnil = ipmem
			continue
		}
	}

	if ipret =="" {
		ipret = ipnil
	}
	return ipret
	//remoteid ,_ := uuid.NewRandom()
	//return remoteid.String()
}
func Print_hostipall()  {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Printf("get_hostip:%s\n", ipnet.IP.String())
					//return ipnet.IP.String()
				}
			}
		}
	}

	//remoteid ,_ := uuid.NewRandom()
	//return remoteid.String()
}