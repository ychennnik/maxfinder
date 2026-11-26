package main

import (
	"fmt"
<<<<<<< HEAD
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
=======
	"math/rand"
	"sync"
	"time"
)

const (
	// SIZE — количество элементов в большом слайсе.
	SIZE = 10_000_000
	// CHUNKS — количество частей, на которые делится слайс в maxChunks.
	CHUNKS = 8
)

// maxRandomValue — верхняя граница случайных чисел (включительно).
// Числа генерируются в диапазоне [1, maxRandomValue], то есть всегда положительные.
const maxRandomValue = 1_000_000_000

// generateRandomElements возвращает слайс из size случайных целых положительных чисел.
// Для size <= 0 возвращает пустой слайс (не nil и без паники).
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(maxRandomValue) + 1
	}

	return elements
}

// maximum возвращает наибольший элемент слайса.
// Для пустого слайса возвращает 0, потому что все допустимые данные положительные.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks делит слайс на CHUNKS частей, ищет максимум каждой части
// в отдельной горутине и возвращает максимум из найденных значений.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	// Если элементов меньше, чем частей, разбиение бессмысленно:
	// размер части был бы равен нулю.
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxima := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		// Если длина не делится на CHUNKS нацело, остаток достаётся последней части,
		// иначе последние элементы слайса остались бы непроверенными.
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			// Каждая горутина пишет только в свою ячейку, поэтому гонки данных нет.
			maxima[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(maxima)
}

func main() {
	fmt.Printf("Генерируем %d целых положительных чисел...\n", SIZE)
	data := generateRandomElements(SIZE)

	start := time.Now()
	singleThreadResult := maximum(data)
	singleThreadMicros := time.Since(start).Microseconds()
	fmt.Printf("Максимальный элемент: %d\nНайден за %d мкс (один поток)\n", singleThreadResult, singleThreadMicros)

	start = time.Now()
	chunksResult := maxChunks(data)
	chunksMicros := time.Since(start).Microseconds()
	fmt.Printf("Максимальный элемент: %d\nНайден за %d мкс (%d горутин)\n", chunksResult, chunksMicros, CHUNKS)
>>>>>>> first-iteration
}
