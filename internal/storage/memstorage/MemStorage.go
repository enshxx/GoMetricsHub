package memstorage

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func New() MemStorage {
	return MemStorage{
		make(map[string]float64),
		make(map[string]int64),
	}
}

func (s *MemStorage) GaugeSet(name string, val float64) {
	s.gauge[name] = val
}

func (s *MemStorage) CounterInc(name string, val int64) {
	if _, exists := s.counter[name]; !exists {
		s.counter[name] = 0
	}
	s.counter[name] += val
}
