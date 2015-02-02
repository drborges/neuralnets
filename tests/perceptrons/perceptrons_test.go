package perceptrons

import (
	"github.com/drborges/neuralnets/perceptrons"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrPerceptron(t *testing.T) {
	assert.Equal(t, 1, perceptrons.Or(1, 1))
	assert.Equal(t, 1, perceptrons.Or(1, 0))
	assert.Equal(t, 1, perceptrons.Or(0, 1))
	assert.Equal(t, 0, perceptrons.Or(0, 0))
}

func TestAndPerceptron(t *testing.T) {
	assert.Equal(t, 1, perceptrons.And(1, 1))
	assert.Equal(t, 0, perceptrons.And(1, 0))
	assert.Equal(t, 0, perceptrons.And(0, 1))
	assert.Equal(t, 0, perceptrons.And(0, 0))
}

func TestXorPerceptron(t *testing.T) {
	assert.Equal(t, 0, perceptrons.Xor(1, 1))
	assert.Equal(t, 1, perceptrons.Xor(1, 0))
	assert.Equal(t, 1, perceptrons.Xor(0, 1))
	assert.Equal(t, 0, perceptrons.Xor(0, 0))
}

func TestNotPerceptron(t *testing.T) {
	assert.Equal(t, 0, perceptrons.Not(1))
	assert.Equal(t, 1, perceptrons.Not(0))
}

func TestNandPerceptron(t *testing.T) {
	assert.Equal(t, 0, perceptrons.Nand(1, 1))
	assert.Equal(t, 1, perceptrons.Nand(1, 0))
	assert.Equal(t, 1, perceptrons.Nand(0, 1))
	assert.Equal(t, 1, perceptrons.Nand(0, 0))
}

func TestNorPerceptron(t *testing.T) {
	assert.Equal(t, 0, perceptrons.Nor(1, 1))
	assert.Equal(t, 0, perceptrons.Nor(1, 0))
	assert.Equal(t, 0, perceptrons.Nor(0, 1))
	assert.Equal(t, 1, perceptrons.Nor(0, 0))
}

func TestXnorPerceptron(t *testing.T) {
	assert.Equal(t, 1, perceptrons.Xnor(1, 1))
	assert.Equal(t, 0, perceptrons.Xnor(1, 0))
	assert.Equal(t, 0, perceptrons.Xnor(0, 1))
	assert.Equal(t, 1, perceptrons.Xnor(0, 0))
}
