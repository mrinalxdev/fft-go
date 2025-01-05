package main

import (
	"math"
	"math/cmplx"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func generateSignal(size int) []complex128 {
	signal := make([]complex128, size)
	for i := 0; i < size; i ++ {
		signal[i] = complex(math.Sin(2*math.Pi*float64(i) / float64(size)), 0)
	}
	return signal
}

func magnitude(c complex128) float64 {
	return cmplx.Abs(c)
}


func plotSignal(signal []complex128, filename string) error {
	points := make(plotter.XYs, len(signal))
	for i := range signal {
		points[i].X = float64(i)
		points[i].Y = real(signal[i])
	}

	p := plot.New()
	p.Title.Text = "Input Signal"
	p.X.Label.Text = "Sample Index"
	p.Y.Label.Text = "Amplitude"

	line, err := plotter.NewLine(points)

	if err != nil {
		return err
	}

	p.Add(line)

	return p.Save(8*vg.Inch, 4*vg.Inch, filename)
}

func plotFFT(fftResults []complex128, filename string) error {
	points := make(plotter.XYs, len(fftResults))
	for i := range fftResults {
		points[i].X = float64(i)
		points[i].Y = magnitude(fftResults[i])
	}

	p := plot.New()
	p.Title.Text = "FFT Magnitude Spectrum"
	p.X.Label.Text = "Frequency Index"
	p.Y.Label.Text = "Magnitude"

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}

	p.Add(line)

	return p.Save(8*vg.Inch, 4*vg.Inch, filename)
}
