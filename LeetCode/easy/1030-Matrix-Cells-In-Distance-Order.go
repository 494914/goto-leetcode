package main

import (
	"fmt"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	res[0] = []int{r0, c0}
	r, c := r0, c0
	for step := 1; step < R*C; {
		r--
		for r < r0 {
			if 0 <= r && c < C {
				res[step] = []int{r, c}
				step++
			}
			r++
			c++
		}
		for c > c0 {
			if r < R && c < C {
				res[step] = []int{r, c}
				step++
			}
			r++
			c--
		}
		for r > r0 {
			if r < R && 0 <= c {
				res[step] = []int{r, c}
				step++
			}
			r--
			c--
		}
		for c < c0 {
			if 0 <= r && 0 <= c {
				res[step] = []int{r, c}
				step++
			}
			r--
			c++
		}
	}
	return res
}

func main() {
	res := allCellsDistOrder(6, 4, 3, 2)
	fmt.Println(res)
}
