package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(divide(10, 0))

	// result, err := divide1(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("result:", result)

	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func divide(l, r int) int {
	return l / r
}

func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return r / l, nil
}

func divide2(l, r int) (result int, err error) { //when panic can't be avoided
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
