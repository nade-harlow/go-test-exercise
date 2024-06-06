package variables

// Swap function swaps the values of x and y
func Swap(x, y any) (any, any) {
	x, y = y, x
	return x, y
}
