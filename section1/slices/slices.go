package slices

// SumEvenNumbers takes a slice of integers and returns the sum of all even numbers
func SumEvenNumbers(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}
