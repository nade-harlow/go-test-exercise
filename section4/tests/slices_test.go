package tests

import (
	"github.com/nade-harlow/go-test-exercise/section1/slices"
	"testing"
)

func TestSumEvenNumbers(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "All even numbers",
			numbers:  []int{2, 4, 6, 8, 10},
			expected: 30,
		},
		{
			name:     "All odd numbers",
			numbers:  []int{1, 3, 5, 7, 9},
			expected: 0,
		},
		{
			name:     "Mixed even and odd numbers",
			numbers:  []int{1, 2, 3, 4, 5, 6},
			expected: 12,
		},
		{
			name:     "No numbers",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "Single even number",
			numbers:  []int{2},
			expected: 2,
		},
		{
			name:     "Single odd number",
			numbers:  []int{3},
			expected: 0,
		},
		{
			name:     "Negative even numbers",
			numbers:  []int{-2, -4, -6, -8, -10},
			expected: -30,
		},
		{
			name:     "Negative and positive even numbers",
			numbers:  []int{-2, -4, 6, 8, -10},
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slices.SumEvenNumbers(tt.numbers)
			if result != tt.expected {
				t.Errorf("sumEvenNumbers(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
