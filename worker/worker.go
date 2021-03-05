package worker

import (
	"fmt"
	"github.com/go-ping/ping"
	"time"
)

type Worker struct {
	startTime time.Time
	endTime   time.Time
	pinger    *ping.Pinger
}

func NewWorker(ip string) *Worker {
	pinger, err := ping.NewPinger(ip)
	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}
	if err != nil {
		return nil
	}
	return &Worker{
		startTime: time.Now(),
		pinger:    pinger,
	}
}

func (worker *Worker) DoPing() {
	err := worker.pinger.Run()
	if err != nil {
		println(err)
	}
	//stats := worker.pinger.Statistics()
}

func (worker *Worker) Stop() {
	worker.pinger.Stop()
}

func (worker *Worker) Summary() string {
	worker.endTime = time.Now()
	stats := worker.pinger.Statistics()

	s1 := fmt.Sprintf("\n--- %s ping statistics ---\n", stats.Addr)
	s2 := fmt.Sprintf("startTime: %s , endTime: %s \n", worker.startTime, worker.endTime)
	s3 := fmt.Sprintf("%d packets transmitted, %d packets received, %v%% packet loss\n",
		stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	s4 := fmt.Sprintf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
		stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	return s1 + s2 + s3 + s4
}
