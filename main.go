package main

import (
	"fmt"
)



func countArray (target int, array []int) (out int) {
	for _,v := range array {
		if v == target {
			out++
		}
	}
	return
}

func sumArray(array []int) (out int) {
	for _,v:=range array {
		out+=v
	}
	return
}

func inventorySteps (limit int) [][]int {

	var outer [][]int
	var counter int
	var base []int
	var storage []int

	for limit > 0 {
		ans := countArray(counter,base) + countArray(counter,storage)
		base = append(base, ans)
		if ans != 0 {
			counter++
			continue
		}
		storage = append(storage, base...)
		outer = append(outer, base)

		counter = 0
		base = []int{}
		limit--
	}
	return outer
} 


func main () {
	var result [][]int = inventorySteps(50)
	for _,v:=range result {
		fmt.Printf("%v -> Sum: %v\n",v,sumArray(v))
	}
}
