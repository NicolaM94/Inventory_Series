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

func inventorySteps (limit int) [][]int {

	var outer [][]int

	var round int = limit
	var counter int
	var base []int
	var storage []int

	for round > 0 {
		fmt.Printf("Counter: %v\n",counter)
		fmt.Printf("Number of %v in base: %v\n",counter,countArray(counter,base))
		fmt.Printf("Number of %v in storage: %v\n",counter,countArray(counter,storage))
		ans := countArray(counter,base) + countArray(counter,storage)
		base = append(base, ans)

		if ans != 0 {
			counter++
		} else {
			storage = append(storage, base...)
			outer = append(outer, base)

			counter = 0
			base = []int{}
			round--
			
		}
	}
	return outer
} 


func main () {
	fmt.Println(inventorySteps(10))

}
