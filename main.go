package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	data := make([]int, size)

	for i := range data {
		data[i] = rnd.Int()
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	max := data[0]

	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	sliceOfMaximums :=make([]int, CHUNKS)
	sizeOneSlice := len(data) / CHUNKS
	max := data[0]

	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		startOfSlice := i*sizeOneSlice
		endOfSlice := i*sizeOneSlice + sizeOneSlice
		oneSlice := data[startOfSlice:endOfSlice]

		if i == CHUNKS - 1 {
			oneSlice = data[startOfSlice:]
		}

		go func(data []int){
			defer wg.Done()
			index := i

			sliceOfMaximums[index] = maximum(oneSlice)
		}(oneSlice)
	}
	wg.Wait()
	max = maximum(sliceOfMaximums)

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
