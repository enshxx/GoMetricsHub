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

func (s *MemStorage) AddGauge(name string, val float64) {
	s.gauge[name] = val
}

func (s *MemStorage) AddCounter(name string, val int64) {
	if _, exists := s.counter[name]; !exists {
		s.counter[name] = 0
	}
	s.counter[name] += val
}

func (s *MemStorage) GetGauge(name string) (float64, bool) {
	val, ok := s.gauge[name]
	return val, ok
}

func (s *MemStorage) GetCounter(name string) (int64, bool) {
	val, ok := s.counter[name]
	return val, ok
}
