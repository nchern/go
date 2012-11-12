package dsp

import (
	//	"log"
	"math"
	"math/rand"
	"time"
)

type ParticleFilter struct {
	FilterBase

	XPlus  [2][]float64
	XMinus [2][]float64

	xPlusCur  int
	xPlusPrev int

	xMinusCur  int
	xMinusPrev int

	ParticleCount int
	ValueRange    float64

	q []float64
}

func (f *ParticleFilter) Init(length int) Filter {
	f.FilterBase.Init(length)

	f.XPlus[0] = make([]float64, f.ParticleCount)
	f.XPlus[1] = make([]float64, f.ParticleCount)
	f.XMinus[0] = make([]float64, f.ParticleCount)
	f.XMinus[1] = make([]float64, f.ParticleCount)

	f.q = make([]float64, f.ParticleCount)

	f.xPlusCur = 0
	f.xPlusPrev = -1

	f.xMinusCur = 0
	f.xMinusPrev = -1

	rand.Seed(time.Now().UnixNano())

	return f
}

func (f *ParticleFilter) particles(x0 float64) []float64 {
	r := NaNArray(f.ParticleCount)
	for i, _ := range r {
		r[i] = x0 + (-1+2*rand.Float64())*f.ValueRange
	}
	return r
}

func (f *ParticleFilter) Step(i int, inputs ...[]float64) error {
	inc := func(i int) int {
		return (i + 1) % 2
	}
	defer func() {
		f.xPlusCur = inc(f.xPlusCur)
		f.xPlusPrev = inc(f.xPlusPrev)
		f.xMinusCur = inc(f.xMinusCur)
		f.xMinusPrev = inc(f.xMinusPrev)
	}()

	x := inputs[0]
	if i == 0 {
		f.XPlus[0] = f.particles(x[0])
		f.XMinus[0] = NaNArray(f.ParticleCount) //x[0]
		return nil
	}
	//f.XMinus[i] = make([]float64, f.ParticleCount)
	//f.XPlus[i] = make([]float64, f.ParticleCount)

	//propagate
	for j := 0; j < f.ParticleCount; j++ {
		f.XMinus[f.xMinusCur][j] = f.XPlus[f.xPlusPrev][j] + (-1+2*rand.Float64())*f.ValueRange //TODO: add propagation noise
	}

	for j := 0; j < f.ParticleCount; j++ {
		f.q[j] = 1 / math.Abs(f.XMinus[f.xMinusCur][j]-x[i])
	}
	f.q = Normalize(f.q)

	//resampling
	for j := 0; j < f.ParticleCount; j++ {
		f.XPlus[f.xPlusCur][j] = f.XMinus[f.xMinusCur][f.choose()] //  from xminus by q WITH REPLACEMENT
	}

	expectedV := 0.
	for _, v := range f.XPlus[f.xPlusCur] {
		expectedV += v
	}
	f.values[i] = expectedV / float64(f.ParticleCount)

	return nil
}

func (f *ParticleFilter) choose() int {
	rnd := rand.Float64()
	s := 0.
	j := 0
	for i, v := range f.q {
		j = i
		s += v
		if s >= rnd {
			break
		}
	}
	return j
}
