package monitor

import (
	"log"
	"os"
)

type Monitor struct {
	logger    *log.Logger
	MonitorCb func() any
}

func (monitor Monitor) Discover() {
	for {
		monitor.MonitorCb()
	}
}

func NewMonitor(cb func() any) *Monitor {
	monitor := Monitor{
		logger:    log.New(os.Stdout, "Monitor ", log.LstdFlags),
		MonitorCb: cb,
	}

	return &monitor
}
