package memstorage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStorage_AddGauge(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name string
		val  float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "base test",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 5,
				},
			},
			args: args{
				name: "cpu_usage",
				val:  5,
			},
			want: 5,
		},
		{
			name: "test empty",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 5,
				},
			},
			args: args{
				name: "cache_size",
				val:  1,
			},
			want: 1,
		},
		{
			name: "test int",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 5,
				},
			},
			args: args{
				name: "cache_size",
				val:  1.0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			s.AddGauge(tt.args.name, tt.args.val)
			res, _ := s.GetGauge(tt.args.name)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMemStorage_AddCounter(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name string
		val  int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "base test",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 5,
				},
			},
			args: args{
				name: "memory_usage",
				val:  5,
			},
			want: 10,
		},
		{
			name: "test empty",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 5,
				},
			},
			args: args{
				name: "cache_size",
				val:  1,
			},
			want: 1,
		},
		{
			name: "test negative",
			fields: fields{
				gauge: map[string]float64{
					"cpu_usage": 4.0,
				},
				counter: map[string]int64{
					"memory_usage": 12,
				},
			},
			args: args{
				name: "memory_usage",
				val:  -10,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			s.AddCounter(tt.args.name, tt.args.val)
			res, _ := s.GetCounter(tt.args.name)
			assert.Equal(t, tt.want, res)
		})
	}
}
