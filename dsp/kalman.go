package dsp

import (
	"math"
)

type KalmanFilter struct {
	FilterBase
	PMinus []float64
	PPlus  []float64
	XMinus []float64
	XPlus  []float64

	K []float64

	Q     float64
	R     float64
	Alpha float64
}

func (f *KalmanFilter) Init(length int) Filter {
	f.FilterBase.Init(length)

	f.PMinus = NaNArray(length)
	f.PPlus = NaNArray(length)
	f.XMinus = NaNArray(length)
	f.XPlus = NaNArray(length)
	f.K = NaNArray(length)

	f.Alpha = 1

	return f
}

func (f *KalmanFilter) Calc(i int, inputs ...[]float64) error {
	x := inputs[0]
	if i == 0 {
		f.XMinus[i] = x[i]
		f.XPlus[i] = x[i]
		return nil
	}

	f.PMinus[i] = math.Pow(f.Alpha, 2)*f.PPlus[i-1] + f.Q
	f.K[i] = f.PMinus[i] / (f.PMinus[i] + f.R)
	f.XMinus[i] = f.XPlus[i-1]
	f.XPlus[i] = f.XMinus[i] + f.K[i]*(x[i]-f.XMinus[i])
	f.PPlus[i] = (1-f.K[i])*f.PMinus[i]*(1-f.K[i]) + f.K[i]*f.K[i]*f.R

	return nil
}

func (f *KalmanFilter) Values() []float64 {
	return f.XPlus[:]
}

type SMA struct {
	FilterBase
	Period int
}

func (f *SMA) Step(i int, inputs ...[]float64) error {
	x := i - f.Period + 1
	if x < 0 {
		return nil
	}
	avg := 0.0
	for _, v := range inputs[0][x : i+1] {
		avg += v
	}
	f.values[i] = avg / float64(f.Period)
	return nil
}
