package main

import (
	"fmt"
	
)


func countMap (number int,data map[int]int) int {
	for k,v := range data {
		if k == number {
			if v == 0 || k == 0 {
				return 1
			}
			return v
		}
	}
	return 0
}

func countNewBase (number int, data []int) (out int) {
	for _,v := range data {
		if v == number {
			out++
		}
	}
	return
}


func inventorySeries (limit int) map[int]int {

	var base map[int]int = make(map[int]int)
	var counter int
	var newBase []int

	for counter <= limit {

		countedMap := countMap(counter,base)
		fmt.Println("CountedMap: ",countedMap)
		//time.Sleep(5* time.Second)
		countedNewBase := countNewBase(counter,newBase)
		fmt.Println("CountedNewBase: ",countedNewBase)
		//time.Sleep(5*time.Second)

		newBase = append(newBase,countedMap+countedNewBase)
		fmt.Println(newBase)
		//time.Sleep(5*time.Second)
		
		if newBase[len(newBase)-1] == 0 {
			for i,_ := range newBase {
				base[i] = countNewBase(base[i],newBase)
			}
			counter = 0
			newBase = newBase[:0]
			fmt.Println(base,"\n")
			continue
		} else {
		fmt.Println(base,"\n")
		counter++
		}
	}
	return base
}






func main () {

	fmt.Println(inventorySeries(2))
	
	
}

