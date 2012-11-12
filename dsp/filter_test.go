package dsp

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
)

func FailOnError(err error) {
	if err == nil {
		return
	}
	log.Fatalf("FATAL: %v", err)
}

func sample(n int, stdev float64) []float64 {
	y := NaNArray(n)
	for i, _ := range y {
		y[i] = float64(i) + rand.NormFloat64()*stdev
	}
	return y
}

func Dump2File(filename string, y []float64, filter, benchmark Filter) {
	f, err := os.Create(filename)
	defer f.Close()
	FailOnError(err)
	for i, yv := range y {
		fmt.Fprintln(f, i, yv, filter.Values()[i], benchmark.Values()[i])
	}
}

func TestKalman(t *testing.T) {

	n := 200
	stdev := 15.0

	p0 := 100000000.0
	kalman := &KalmanFilter{}
	kalman.Init(n)
	kalman.PMinus[0] = p0
	kalman.PPlus[0] = p0
	kalman.Q = stdev / 2
	kalman.R = stdev * 1
	kalman.Alpha = 1.1

	y := sample(n, stdev)

	FailOnError(BatchFilter(kalman, y))

	sma := &SMA{Period: 2}
	sma.Init(len(y))

	FailOnError(BatchFilter(sma, y))

	Dump2File("/tmp/kalman.txt", y, kalman, sma)

}

func TestParticle(t *testing.T) {

	n := 100

	filter := &ParticleFilter{ParticleCount: 3000, ValueRange: 30}
	filter.Init(n)

	stdev := 30.0

	y := sample(n, stdev)

	FailOnError(BatchFilter(filter, y))

	sma := &SMA{Period: 3}
	sma.Init(len(y))

	FailOnError(BatchFilter(sma, y))

	Dump2File("/tmp/filter.txt", y, filter, sma)
}
