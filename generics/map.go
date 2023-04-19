package main

import "fmt"

func main() {
	testScores := map[string]float32{
		"Harry":    87.3,
		"Hermione": 105,
		"Ronald":   63.5,
		"Neville":  27,
	}

	c := clone(testScores)

	fmt.Println(c)
}

func clone[K comparable, V any](m map[K]V) map[K]V { //Key in map must be comparable
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func clone_strong(m map[string]float64) map[string]float64 {
	result := make(map[string]float64, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
