package main

import "fmt"

func main() {
	const a = 3
	const b float32 = 3
	var f32 float32 = b
	var f64 float64 = float64(b)

	fmt.Println(f32, f64)

	const c = iota //iota is related to position in group, here 0
	fmt.Println(c)

	const (
		d = 2 * 5
		e             // e is in this case carried over, so it's like 2 * 5
		f = iota      //here iota is 2 (3rd in the group)
		g             // ioata is carried on, so it's 3
		h = 10 * iota //40
	)
	fmt.Println(d, e, f, g, h)

}
