package main

<<<<<<< HEAD
// Пишите тесты в этом файле
=======
import (
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{name: "нулевой размер", size: 0, wantSize: 0},
		{name: "отрицательный размер", size: -5, wantSize: 0},
		{name: "один элемент", size: 1, wantSize: 1},
		{name: "обычный размер", size: 1000, wantSize: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if got == nil {
				t.Fatal("ожидался непустой указатель на слайс (пустой слайс допустим), получен nil")
			}
			if len(got) != tt.wantSize {
				t.Fatalf("длина слайса = %d, ожидалось %d", len(got), tt.wantSize)
			}
			for i, value := range got {
				if value <= 0 || value > maxRandomValue {
					t.Fatalf("элемент %d = %d вне диапазона [1, %d]", i, value, maxRandomValue)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 1, 2, 3}, want: 9},
		{name: "максимум в конце", data: []int{1, 2, 3, 9}, want: 9},
		{name: "максимум в середине", data: []int{1, 9, 3}, want: 9},
		{name: "все элементы равны", data: []int{5, 5, 5}, want: 5},
		{name: "несколько одинаковых максимумов", data: []int{7, 3, 7}, want: 7},
		{name: "отрицательные числа", data: []int{-3, -1, -2}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	// Размеры вокруг границ: меньше, равно и больше CHUNKS, кратные и некратные CHUNKS.
	sizes := []int{0, 1, 2, CHUNKS - 1, CHUNKS, CHUNKS + 1, 2*CHUNKS - 1, 2 * CHUNKS, 1000, 1003}

	for _, size := range sizes {
		// Максимум ставим в каждую позицию слайса: так проверяются все части,
		// включая остаток в конце слайса, который легко потерять при делении на части.
		for position := 0; position < size; position++ {
			data := make([]int, size)
			for i := range data {
				data[i] = 1
			}
			data[position] = 100

			if got := maxChunks(data); got != 100 {
				t.Fatalf("размер %d, максимум на позиции %d: maxChunks = %d, ожидалось 100", size, position, got)
			}
		}
	}

	t.Run("пустой слайс", func(t *testing.T) {
		if got := maxChunks([]int{}); got != 0 {
			t.Fatalf("maxChunks(пустой) = %d, ожидалось 0", got)
		}
		if got := maxChunks(nil); got != 0 {
			t.Fatalf("maxChunks(nil) = %d, ожидалось 0", got)
		}
	})

	t.Run("совпадает с maximum на случайных данных", func(t *testing.T) {
		random := rand.New(rand.NewSource(1))
		for round := 0; round < 100; round++ {
			data := make([]int, random.Intn(5000))
			for i := range data {
				data[i] = random.Intn(1_000_000) + 1
			}

			if got, want := maxChunks(data), maximum(data); got != want {
				t.Fatalf("размер %d: maxChunks = %d, maximum = %d", len(data), got, want)
			}
		}
	})
}

func BenchmarkMaximum(b *testing.B) {
	data := generateRandomElements(SIZE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximum(data)
	}
}

func BenchmarkMaxChunks(b *testing.B) {
	data := generateRandomElements(SIZE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxChunks(data)
	}
}
>>>>>>> first-iteration
