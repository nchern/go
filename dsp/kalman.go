package dsp

import (
	//	"log"
	"math"
)

type KalmanFilter struct {
	FilterBase
	PMinus float64
	XMinus float64
	PPlus  []float64
	PPlus0 float64
	XPlus  []float64

	K float64

	Q     float64
	R     float64
	Alpha float64
}

func (f *KalmanFilter) Init(length int) Filter {
	f.FilterBase.Init(length)

	f.PPlus = NaNArray(length)
	f.PPlus[0] = f.PPlus0
	f.XPlus = NaNArray(length)

	return f
}

func (f *KalmanFilter) Calc(i int, inputs ...[]float64) error {
	x := inputs[0]
	if i == 0 {
		f.XMinus = x[i]
		f.XPlus[i] = x[i]
		return nil
	}

	f.PMinus = math.Pow(f.Alpha, 2)*f.PPlus[i-1] + f.Q
	f.K = f.PMinus / (f.PMinus + f.R)
	f.XMinus = f.XPlus[i-1]
	f.XPlus[i] = f.XMinus + f.K*(x[i]-f.XMinus)
	f.PPlus[i] = (1-f.K)*f.PMinus*(1-f.K) + f.K*f.K*f.R
	return nil
}

func (f *KalmanFilter) Values() []float64 {
	return f.XPlus[:]
}

func (f *KalmanFilter) String() string {
	return "kalman"
}

type SMA struct {
	FilterBase
	Period int
}

func (f *SMA) Calc(i int, inputs ...[]float64) error {
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
