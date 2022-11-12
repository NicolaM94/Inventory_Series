package main

import "fmt"

var BASE map[int]int = make(map[int]int)

func countMapOccurencies(n int) (out int) {
	for k, v := range BASE {
		if k == n {
			out += n
		}
		if BASE[k] == n {
			out += BASE[v]
		}
	}
	return
}

func main() {

	BASE[0] = 0

	//T(1)
	BASE[0] = countMapOccurencies(0)
	BASE[1] = countMapOccurencies(1)

	fmt.Println(BASE)

	//T(2)
	BASE[0] = countMapOccurencies(0)
	BASE[1] = countMapOccurencies(1)
	BASE[2] = countMapOccurencies(2)

}
