package lib_test

import (
	"time"
	"sync"
	"errors"
)

type MockStatter struct {
	sync.Mutex
	ContStats map[string]int64
	GaugeStats map[string]int64
	TimingStats map[string][]int64
	DurationStats map[string][]time.Duration
	RawStats map[string]string
	Closed bool
	Err error
	Prefix string
}

func (s *MockStatter) Inc(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Inc called on closed statter")
		return s.Err
	}
	if s.ContStats == nil {
		s.ContStats = map[string]int64{}
	}
	if _, ok := s.ContStats[stat]; !ok {
		s.ContStats[stat] = value
	} else {
		s.ContStats[stat] += value
	}
	return nil
}

func (s *MockStatter) Dec(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Dec called on closed statter")
		return s.Err
	}
	if s.ContStats == nil {
		s.ContStats = map[string]int64{}
	}
	if _, ok := s.ContStats[stat]; !ok {
		s.ContStats[stat] = 0 - value
	} else {
		s.ContStats[stat] -= value
	}
	return nil
}

func (s *MockStatter) Gauge(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Gauge called on closed statter")
		return s.Err
	}
	if s.GaugeStats == nil {
		s.GaugeStats = map[string]int64{}
	}
	s.GaugeStats[stat] = value
	return nil
}

func (s *MockStatter) GaugeDelta(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("GaugeDelta called on closed statter")
		return s.Err
	}
	if s.GaugeStats == nil {
		s.GaugeStats = map[string]int64{}
	}
	if _, ok := s.GaugeStats[stat]; !ok {
		s.GaugeStats[stat] = value
	} else {
		s.GaugeStats[stat] += value
	}
	return nil
}

func (s *MockStatter) Timing(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Timing called on closed statter")
		return s.Err
	}
	if s.TimingStats == nil {
		s.TimingStats = map[string][]int64{}
	}
	if _, ok := s.TimingStats[stat]; !ok {
		s.TimingStats[stat] = []int64{value}
	} else {
		s.TimingStats[stat] = append(s.TimingStats[stat], value)
	}
	return nil
}

func (s *MockStatter) TimingDuration(stat string, value time.Duration, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("TimingDuration called on closed statter")
		return s.Err
	}
	if s.DurationStats == nil {
		s.DurationStats = map[string][]time.Duration{}
	}
	if _, ok := s.DurationStats[stat]; !ok {
		s.DurationStats[stat] = []time.Duration{value}
	} else {
		s.DurationStats[stat] = append(s.DurationStats[stat], value)
	}
	return nil
}

func (s *MockStatter) Set(stat string, value string, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Set called on closed statter")
		return s.Err
	}
	if s.RawStats == nil {
		s.RawStats = map[string]string{}
	}
	s.RawStats[stat] = value
	return nil
}

func (s *MockStatter) SetInt(stat string, value int64, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("SetInt called on closed statter")
		return s.Err
	}
	if s.ContStats == nil {
		s.ContStats = map[string]int64{}
	}
	s.ContStats[stat] = value
	return nil
}

func (s *MockStatter) Raw(stat string, value string, rate float32) error {
	s.Lock()
	defer s.Unlock() 
	if s.Closed {
		s.Err = errors.New("Raw called on closed statter")
		return s.Err
	}
	if s.RawStats == nil {
		s.RawStats = map[string]string{}
	}
	s.RawStats[stat] = value
	return nil
}

func (s *MockStatter) SetPrefix(prefix string) {
	s.Lock()
	defer s.Unlock()
	s.Prefix = prefix
	return
}

func (s *MockStatter) Close() error {
	if s.Closed {
		s.Err = errors.New("SetPrefix called on closed statter")
		return s.Err
	}
	s.Closed = true
	return nil
}