package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHaversineDistance(t *testing.T) {
	expected := float64(1)
	actual := haversineDistance(35.0237056, 135.7545413, 35.0237009, 135.7579916)
	assert.LessOrEqual(t, actual, expected)
	expected = float64(0)
	assert.GreaterOrEqual(t, actual, expected)

	/*actual := haversineDistance(35.0237056, 135.7545413, 35.0237009, 135.7579916)

	expected := float64(0)
	assert.Equal(t, expected, actual)

	actual := haversineDistance(35.0237056, 135.7545413, 35.0237009, 135.7579916)

	expected := float64(0)
	assert.Equal(t, expected, actual)*/
}
