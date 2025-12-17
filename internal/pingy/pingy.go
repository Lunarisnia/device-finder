package pingy

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func Ping(addr string) (int, error) {
	pinger, err := probing.NewPinger(addr)
	if err != nil {
		return 0, err
	}
	pinger.Timeout = 5 * time.Second
	pinger.Count = 5
	err = pinger.Run()
	if err != nil {
		return 0, err
	}
	stats := pinger.Statistics()
	return stats.PacketsRecv, nil
}
