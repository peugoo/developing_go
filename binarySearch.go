package main

import (
	"fmt"
	"sort"
)

func binary_search(n int, mass []int) int {

	sort.Slice(mass, func(i, j int) bool {
		return mass[i] < mass[j]
	}) 

	splitSlice := func(k, n int, sl []int) ([]int, int) {
		if k <= sl[len(sl) / 2 - 1] {
			return sl[:len(sl) / 2], n
		} else {
			return sl[len(sl) / 2:], n + len(sl) / 2
		}
	}

	ind := 0
	for {
		mass, ind = splitSlice(n, ind, mass)
		if len(mass) == 1 {
			if mass[0] == n {
				return ind
			} else {
				return -1
			}
		}
	}

}

func main(){
	var a []int = []int{}
	for i := 500; i != -63; i-- {
		a = append(a, i)
	}
	a = append(a, -6)
	for i := 100000; i != 100000000; i++ {
		a = append(a, i)
	}
	fmt.Print(binary_search(1000412, a))
}
