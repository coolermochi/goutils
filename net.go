package goutils

import (
	"io"
	"net"
	"time"
)

type network struct {}

var Net = network{}

func (network) Ping(ip string) (bool, error) {
	one := []byte{}
	conn, err := net.DialTimeout("tcp", ip, time.Second*2)
	if err == nil {
		conn.SetReadDeadline(time.Now())
		if _, err := conn.Read(one); err == io.EOF {
			conn.Close()
			conn = nil
			return false, err
		}
		return true, nil
	}
	return false, err
}
