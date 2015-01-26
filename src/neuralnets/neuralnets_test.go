package neuralnets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrPerceptron(t *testing.T) {
	assert.Equal(t, 1, OrPerceptron(1, 1))
	assert.Equal(t, 1, OrPerceptron(1, 0))
	assert.Equal(t, 1, OrPerceptron(0, 1))
	assert.Equal(t, 0, OrPerceptron(0, 0))
}

func TestAndPerceptron(t *testing.T) {
	assert.Equal(t, 1, AndPerceptron(1, 1))
	assert.Equal(t, 0, AndPerceptron(1, 0))
	assert.Equal(t, 0, AndPerceptron(0, 1))
	assert.Equal(t, 0, AndPerceptron(0, 0))
}

func TestXorPerceptron(t *testing.T) {
	assert.Equal(t, 0, XorPerceptron(1, 1))
	assert.Equal(t, 1, XorPerceptron(1, 0))
	assert.Equal(t, 1, XorPerceptron(0, 1))
	assert.Equal(t, 0, XorPerceptron(0, 0))
}
