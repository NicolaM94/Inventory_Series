package main

import (
	"inventory_series/utilities"
	"log"
	"os"
)


func inventorySteps (limit int) [][]int {

	var outer [][]int
	var counter int
	var base []int
	var storage []int

	for limit > 0 {
		ans := utilities.CountArray(counter,base) + utilities.CountArray(counter,storage)
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
	var result [][]int = inventorySteps(500)
	file,err := os.Create("/home/nicola/Desktop/tet.txt")
	if err != nil {
		panic("Cannot open target file")
	}

	for _,v:=range result {
		//fmt.Printf("Writing %v",utilities.StringArray(v))
		_,err := file.Write(utilities.StringArray(v))

		if err != nil {
			log.Fatal(err)
		}
	}
	closeErr := file.Close()

	if closeErr != nil {
		log.Fatal(closeErr)
	}
	
}
