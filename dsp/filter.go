package dsp

import (
	"fmt"
	"math"
)

type Filter interface {
	Values() []float64
	Init(int) Filter
	Calc(int, ...[]float64) error
}

type FilterBase struct {
	values []float64
}

func (f *FilterBase) Init(length int) Filter {
	f.values = NaNArray(length) //make([]float64, length)
	return f
}

func (f *FilterBase) Values() []float64 {
	return f.values
}

func (f *FilterBase) Calc(i int, inputs ...[]float64) error {
	return fmt.Errorf("Not implemented")
}

func BatchFilter(f Filter, inputs ...[]float64) error {
	length := len(inputs[0])
	for i := 0; i < length; i++ {
		e := f.Calc(i, inputs...)
		if e != nil {
			return e
		}
	}
	return nil
}

func NaNArray(length int) []float64 {
	r := make([]float64, length)
	for i := range r {
		r[i] = math.NaN()
	}
	return r
}

func Normalize(arr []float64) []float64 {
	r := make([]float64, len(arr))
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	for i, v := range arr {
		r[i] = v / sum
	}
	return r
}
