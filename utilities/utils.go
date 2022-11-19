package utilities

import "fmt"


func CountArray (target int, array []int) (out int) {
	for _,v := range array {
		if v == target {
			out++
		}
	}
	return
}

func SumArray(array []int) (out int) {
	for _,v:=range array {
		out+=v
	}
	return
}

func AverageArray (array []int) int {
	return SumArray(array)/len(array)
}

func StringArray (array []int) []byte {
	var out string
	for _,v := range array {
		out += fmt.Sprint(v)
		out += " "
	}
	out += "\n"
	return []byte(out)
}