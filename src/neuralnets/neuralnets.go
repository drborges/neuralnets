package neuralnets

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

func OrPerceptron(x1, x2 float64) int {
	activation := 1.0
	return Perceptron([]float64{x1, x2}, []float64{1, 1}, activation)
}

func AndPerceptron(x1, x2 float64) int {
	activation := 2.0
	return Perceptron([]float64{x1, x2}, []float64{1, 1}, activation)
}

func XorPerceptron(x1, x2 float64) int {
	activation := 1.0
	and := float64(AndPerceptron(x1, x2))
	return Perceptron([]float64{x1, x2, and}, []float64{1, 1, -2}, activation)
}
