package gosrc

import (
	"math/rand"
	"net"
	"strconv"
	"time"
)

func CheckTcpConnect(host string, port string) (err error) {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return
	}
	if conn != nil {
		defer conn.Close()
		return
	}
	return
}

func GetUnUsePort() uint32 {
	for i := 0; i < 10; i++ {
		if i < 1024 {
			continue
		}
		var newPort = rand.Intn(65535)
		if err := CheckTcpConnect("127.0.0.1", strconv.Itoa(newPort)); nil != err {
			return uint32(newPort)
		}
	}
	return 0
}
