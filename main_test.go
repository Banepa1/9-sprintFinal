package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {	
	assert.Nil(t, generateRandomElements(0))
	assert.Nil(t, generateRandomElements(-1))

	len := 6
	data := generateRandomElements(len)
	assert.Len(t, data, len)
}

func TestMaximum(t *testing.T) {
	var nilSlice []int
	assert.Zero(t, maximum(nilSlice))

	zeroLenSlice := make([]int, 0)
	assert.Zero(t, maximum(zeroLenSlice))

	max := 6

	var oneElementSlice = []int{max}
	assert.Equal(t, max, maximum(oneElementSlice))

	var firstMaxSlice = []int{max, 4, 3, 2}
	assert.Equal(t, max, maximum(firstMaxSlice))

	var endMaxSlice = []int{3, 4, 3, max}
	assert.Equal(t, max, maximum(endMaxSlice))

	var equalElementsSlice = []int{max, max, max, max, max, max, max, max}
	assert.Equal(t, max, maximum(equalElementsSlice))
}

func TestMaxChunks(t *testing.T) {
	var nilSlice []int
	assert.Zero(t, maxChunks(nilSlice))

	zeroLenSlice := make([]int, 0)
	assert.Zero(t, maxChunks(zeroLenSlice))

	max := 300

	var oneElementSlice = []int{max}
	assert.Equal(t, max, maxChunks(oneElementSlice))

	var lessEightSlice = []int{6, 10, 200, 100, max}
	assert.Equal(t, max, maxChunks(lessEightSlice))

	var eightSlice = []int{7, 4, 3, 3, max, 23, 254, 10}
	assert.Equal(t, max, maxChunks(eightSlice))

	var notDivisibleBy8Slice = []int{7, 4, 3, 3, max, 23, 254, 10, 15}
	assert.Equal(t, max, maxChunks(notDivisibleBy8Slice))

	var equalElementsSlice = []int{max, max, max, max, max, max, max, max, max, max ,max}
	assert.Equal(t, max, maxChunks(equalElementsSlice))

	bigSlice := make([]int, 1_000_000)
	for i := range bigSlice {
		bigSlice[i] = i
	}

	assert.Equal(t, 999_999, maxChunks(bigSlice))
}