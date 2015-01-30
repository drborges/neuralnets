package perceptrons

func Perceptron(inputs []float64, weights []float64, activation float64) int {
	var sum = 0.0
	for i, input := range inputs {
		sum += input * weights[i]
	}

	if sum >= activation {
		return 1
	}

	return 0
}

func Or(x1, x2 float64) int {
	activation := 1.0
	return Perceptron([]float64{x1, x2}, []float64{1, 1}, activation)
}

func And(x1, x2 float64) int {
	activation := 2.0
	return Perceptron([]float64{x1, x2}, []float64{1, 1}, activation)
}

func Xor(x1, x2 float64) int {
	activation := 1.0
	and := float64(And(x1, x2))
	return Perceptron([]float64{x1, x2, and}, []float64{1, 1, -2}, activation)
}

func Not(x1 float64) int {
	return Xor(x1, 1)
}