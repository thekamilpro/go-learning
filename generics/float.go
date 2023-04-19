package main

import "fmt"

func main() {
	testScores := []float32{ //change float32 to float64 and it still works
		87.3,
		105,
		63.5,
		27,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)
}

func clone[V any](s []V) []V { //V is generic, and is constraint
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
