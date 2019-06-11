package main

import (
	"os"
	"fmt"
	"math"
	"bufio"
	"strconv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type History struct {
	Label        string
	Measurements []float64
}

type Neuron struct {
	Observations []History
	Weights      []float64
	Encoding     map[string]float64
	Decoding 	 map[float64]string
	LearningRate float64
}

func (n *Neuron) EncodeLabels() {
	/* Turn any string labels into numbers for our model. */
	
	n.Encoding = make(map[string]float64)
	for _, o := range n.Observations {
		if _, ok := n.Encoding[o.Label]; !ok {
			n.Encoding[o.Label] = float64(len(n.Encoding))
		}
	}
}

func (n *Neuron) DecodeLabels() {
	/* Create a number to string text for human readability */
	
	n.Decoding = make(map[float64]string)
	for label, code := range n.Encoding {
		n.Decoding[code] = label
	}
}

func (n *Neuron) Decode(pred float64) string {
	/* Turn a decimal number to nearest categorical number in decoding map */
	
	var minimum float64 = 1000000
	var decoded string = ""

	for label, code := range n.Encoding {

		if minimum > math.Abs(pred - code) {
			minimum = math.Abs(pred - code)
			decoded = label
		}

	
	}

	return decoded

}

func (n *Neuron) PrepareWeights() {
	n.Weights = make([]float64, len(n.Observations[0].Measurements)+1)
}

func (n *Neuron) Sigmoid(input float64) float64 {
	return 1 / (1 + math.Exp(-1*input))
}

func (n *Neuron) Process() float64 {
	costs := 0.0
	for _, observed := range n.Observations {

		predicted := n.Predict(observed.Measurements...)
		target := n.Encoding[observed.Label]
		error_delta := predicted - target
		predicted_delta := predicted * ( 1 - predicted )

		// Update weights.
		for i:=0; i<len(n.Weights)-1; i++ {

			n.Weights[i] = n.Weights[i] - n.LearningRate * error_delta * predicted_delta * observed.Measurements[i]

		}

		// Update bias weight.
		n.Weights[len(n.Weights)-1] = n.Weights[len(n.Weights)-1] - n.LearningRate * error_delta * predicted_delta

		costs += (error_delta * error_delta)/2.0
	}

	return costs
}

func (n *Neuron) Train(iterations int) {
	var costs []float64

	for iteration:=0; iteration<iterations; iteration++ {

		costs = append(costs, n.Process())

	}

	n.Plot(costs)
	
}

func (n *Neuron) Learn(iterations int) {
	n.EncodeLabels()
	n.PrepareWeights()
	n.LearningRate = 0.01
	n.Train(iterations)
}

func (n *Neuron) Predict(new_measurement ...float64) float64 {

	// Label = sigmoid( Length * w1 + Width * w2 + bias )
	return n.Sigmoid(new_measurement[0]*n.Weights[0] + new_measurement[1]*n.Weights[1] + n.Weights[2])
}

func (n *Neuron) Plot(input []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Machine Learning Cost"
	p.X.Label.Text = "Iteration"
	p.Y.Label.Text = "Cost"

	pts := make(plotter.XYs, len(input))

	for index, value := range input {
		pts[index].X, pts[index].Y = float64(index), value
	}

	err = plotutil.AddLinePoints(p,
		"Cost", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "costs.png"); err != nil {
		panic(err)
	}
}

func main() {

	var brain Neuron = Neuron{
		Observations: []History{
			History{"red", []float64{3, 1.5}},
			History{"blue", []float64{2, 1}},
			History{"red", []float64{4, 1.5}},
			History{"blue", []float64{3, 1}},
			History{"red", []float64{3.5, 0.5}},
			History{"blue", []float64{2, 0.5}},
			History{"red", []float64{5.5, 1}},
			History{"blue", []float64{1, 1}},
		},
	}

	brain.EncodeLabels()
	fmt.Println("BRAIN LABEL ENCODING: ", brain.Encoding)
	brain.DecodeLabels()
	fmt.Println("BRAIN LABEL ENCODING: ", brain.Decoding)

	brain.PrepareWeights()
	fmt.Println("BRAIN CONNECTION WEIGHTS: ", brain.Weights)

	brain.LearningRate = 0.01
	fmt.Println("LEARNING RATE SET TO: ", brain.LearningRate)

	fmt.Println("STARTING 1 MILLION TRAINING ITERATIONS")
	brain.Train(1000000)
	fmt.Println("TRAINING COMPLETE")


	scanner := bufio.NewScanner(os.Stdin)

	for {
		
		fmt.Println("What length?")
		scanner.Scan()
		length_string := scanner.Text()
		length, _ := strconv.ParseFloat(length_string, 32)

		fmt.Println("What Width?")
		scanner.Scan()
		width_string := scanner.Text()
		width, _ := strconv.ParseFloat(width_string, 32)

		fmt.Println(brain.Decode(brain.Predict(length, width)))
		fmt.Println()
	}

}
