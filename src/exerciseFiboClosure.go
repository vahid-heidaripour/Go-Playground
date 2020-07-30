package main

import "fmt"

func fibonacci() func() int {
	v2, v1 := 0, 1
	return func() int {
		ret := v2
		v2, v1 = v1, v1+ret

		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
