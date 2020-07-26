package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i := range slice {
		slice[i] = make([]uint8, dx)
	}

	for i := range slice {
		for j := range slice[i] {
			slice[i][j] = uint8((i ^ j) / 2)
			//slice[i][j] = uint8((i + j) / 2)
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
