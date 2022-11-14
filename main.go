package main

import "fmt"

var KEYS []int
var VALUES []int

func contains (number int, array []int) (out int) {
	for _,v := range array {
		if v == number {
			out++
		}
	}
	return
}


func main () {
	
	var base []int = []int{0}
	var counter int = 0
	var limit int = 5
	var newBase = base

	for counter <= limit {
		var temp int = contains(counter,base)
		newBase = append(newBase, contains(counter,base))
		if temp == 0 {
			base = newBase
			counter = 0
			newBase = []int{}
			fmt.Println(base)
			continue
		} else {
			counter++
		}
	}
}

