package main

func toHeap() *int {
	var x int
	return &x
}

func toStack() int {
	x := new(int)
	*x = 1
	return *x
}

/**
 *
 */
func main() {
	var a string
}
