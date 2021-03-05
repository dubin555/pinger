package worker

import (
	"fmt"
	"sync"
	"time"
)

type Master struct {
	workers map[string]*Worker
	summary map[string]string
	lock    sync.Mutex
	ctxs    map[string]chan struct{}// context to stop the refresh goroutine
}

func NewMaster() *Master {
	master := &Master{
		workers: make(map[string]*Worker),
		summary: make(map[string]string),
		lock:    sync.Mutex{},
		ctxs:    make(map[string]chan struct{}),
	}
	return master
}

func (master *Master) Start(address string) {
	master.lock.Lock()
	defer master.lock.Unlock()
	_, ok := master.workers[address]
	if ok {
		return
	}
	master.workers[address] = NewWorker(address)
	go master.workers[address].DoPing()
	master.summary[address] = ""
	master.ctxs[address] = make(chan struct{})

	// launch a ticker
	// todo change ticker interval to configurable
	go func(address string) {
		fmt.Printf("%s refresh goroutine start\n", address)
		d := time.Minute * 1
		t := time.NewTicker(d)
		for {
			select {
			case <-t.C:
				master.refresh(address)
			case <-master.ctxs[address]:
				t.Stop()
				fmt.Printf("%s refresh goroutine exit\n", address)
				break
			}
		}
	}(address)
}

func (master *Master) refresh(address string) {
	master.lock.Lock()
	defer master.lock.Unlock()
	summary := master.workers[address].Summary()
	master.summary[address] = master.summary[address] + summary
	master.workers[address].Stop()
	master.workers[address] = NewWorker(address)
	go master.workers[address].DoPing()
}

func (master *Master) Stop(address string) string {
	master.lock.Lock()
	defer master.lock.Unlock()
	_, ok := master.workers[address]
	if !ok {
		return ""
	}
	summary := master.summary[address] + master.workers[address].Summary()
	master.workers[address].Stop()
	// send signal to kill the refresh goroutine
	master.ctxs[address] <- struct{}{}
	delete(master.workers, address)
	delete(master.summary, address)
	return summary
}

func (master *Master) Summary(address string) string {
	return master.summary[address]
}
