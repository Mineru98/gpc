package network

import (
	"fmt"
	"net"
	"os"
	"time"
)

func tryPort(network string, seq uint64, timeout time.Duration) {
	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", network, timeout)
	endTime := time.Now()
	if err != nil {
		os.Stdout.Write([]byte(startTime.Format("[2006-01-02T15:04:05]:") + " connection failed\n"))
	} else {
		defer conn.Close()
		var t = float64(endTime.Sub(startTime)) / float64(time.Millisecond)
		os.Stdout.Write([]byte(startTime.Format("[2006-01-02T15:04:05]:") + fmt.Sprintf(" addr=%s seq=%d time=%4.2fms\n", conn.RemoteAddr().String(), seq, t)))
	}
}

func ping(ipArg string, timeoutArg int, periodArg int) {
	var seqNumber uint64 = 0
	var network = fmt.Sprintf("%s:%d", ipArg, 80)
	var timeout = time.Duration(timeoutArg) * time.Millisecond
	var period = time.Duration(periodArg) * time.Millisecond

	if timeout >= period {
		fmt.Print(fmt.Errorf("timeout should be less than period"))
		os.Exit(1)
	}

	ticker := time.NewTicker(period)
	quit := make(chan interface{})

	for ; ; seqNumber++ {
		select {
		case <-ticker.C:
			tryPort(network, seqNumber, timeout)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
