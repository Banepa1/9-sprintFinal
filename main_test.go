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
	max := 6

	test := []struct{
		name string
		input  []int
		expect int
	}{
		{"nilslice", nilSlice, 0},
		{"emptySlice", []int{} , 0},
		{"oneElementSlice", []int{max}, max},
		{"firstElement", []int{max, 4, 3, 2}, max},
		{"lastElemt", []int{3, 4, 3, max}, max},
		{"equalElements", []int{max, max, max, max, max, max, max, max}, max},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := maximum(v.input)
			assert.Equal(t, v.expect, got, "%s: got %v, want %v", v.name, got, v.expect)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	var nilSlice []int
	max := 300
	bigSlice := make([]int, 1_000_000)
	for i := range bigSlice {
		bigSlice[i] = i
	}

	test := []struct{
		name string
		input  []int
		expect int
	}{
		{"nilslice", nilSlice, 0},
		{"emptySlice", []int{} , 0},
		{"oneElementSlice", []int{max}, max},
		{"lessEightSlice", []int{6, 10, 200, 100, max}, max},
		{"eightSlice", []int{7, 4, 3, 3, max, 23, 254, 10}, max},
		{"notDivisibleBy8Slice", []int{7, 4, 3, 3, max, 23, 254, 10, 15}, max},
		{"firstElement", []int{max, 4, 3, 2, 100, 200, 150, 120, 7, 15, 23, 45}, max},
		{"lastElemt", []int{4, 3, 2, 100, 200, 150, 120, 7, 15, 23, 45, max}, max},
		{"equalElements", []int{max, max, max, max, max, max, max, max, max, max ,max}, max},
		{"bigSlice", bigSlice, 999_999},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			got := maxChunks(v.input)
			assert.Equal(t, v.expect, got, "%s: got %v, want %v", v.name, got, v.expect)
		})
	}
}