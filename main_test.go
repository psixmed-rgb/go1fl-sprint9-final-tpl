package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantSize int
	}{
		{
			name:     "нулевой размер",
			size:     0,
			wantSize: 0,
		},
		{
			name:     "отрицательный размер",
			size:     -5,
			wantSize: 0,
		},
		{
			name:     "один элемент",
			size:     1,
			wantSize: 1,
		},
		{
			name:     "обычный размер",
			size:     100,
			wantSize: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantSize {
				t.Errorf("получена длина %d, ожидалась %d", len(got), tt.wantSize)
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
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "nil слайс",
			data: nil,
			want: 0,
		},
		{
			name: "один элемент",
			data: []int{42},
			want: 42,
		},
		{
			name: "несколько элементов",
			data: []int{10, 55, 3, 99, 24, 1},
			want: 99,
		},
		{
			name: "одинаковые элементы",
			data: []int{7, 7, 7, 7},
			want: 7,
		},
		{
			name: "максимум в начале",
			data: []int{500, 100, 200, 50},
			want: 500,
		},
		{
			name: "максимум в конце",
			data: []int{10, 20, 30, 400},
			want: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}