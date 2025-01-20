package statistics

import "sync"

type statistics struct {
	vmCount               int
	totalRequests         int
	totalRequestTimeCount int64
	mut                   sync.RWMutex
}

type Statistics interface {
	AddStats(nanoSeconds int64)
	GetStats() *stats
}

func NewApiStatistics(vmCount int) Statistics {
	return &statistics{vmCount: vmCount}
}

func (s *statistics) AddStats(nanoSeconds int64) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.totalRequests++
	s.totalRequestTimeCount += nanoSeconds
}

type stats struct {
	VMCount            int     `json:"vm_count"`
	RequestCount       int     `json:"request_count"`
	AverageRequestTime float64 `json:"average_request_time"`
}

func (s *statistics) GetStats() *stats {
	s.mut.RLock()
	defer s.mut.RUnlock()
	res := &stats{VMCount: s.vmCount, RequestCount: s.totalRequests}
	if s.totalRequests != 0 {
		res.AverageRequestTime = float64(s.totalRequestTimeCount) / float64(s.totalRequests) / 1e6
	}
	return res
}
