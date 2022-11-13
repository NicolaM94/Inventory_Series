package main

import "fmt"

type Rectangle struct {
	a []int
	b []int
	c []int
	d []int
	points [][]int
}

func calculatePoints (A Rectangle) Rectangle {

	A.points = append(A.points, A.a)
	A.points = append(A.points, A.b)

	for i:= A.a[0];i<A.b[0];i++ {
		A.points = append(A.points,[]int{A.a[0]+1,A.a[1]})
	}

	for i,v := range A.points {
		new := v
		new[1] = i+2
		if new[1] > A.d[1] {break}
		A.points = append(A.points, new)
	}
	
	return A
}

func countDuplicates (pointsA,pointsB [][]int) (out int) {
	for _,v:= range pointsA {
		for _,b := range pointsB {
			if v == b {
				//TODO: Fix here
				out
			}
		}
	}
}

func main() {

	var A Rectangle = Rectangle {
		a: []int{1,4},
		b: []int{4,4},
		c: []int{4,7},
		d: []int{1,7},
		points: [][]int{},
	}

	var B Rectangle = Rectangle {
		a: []int{0,5},
		b: []int{4,5},
		c: []int{4,8},
		d: []int{3,8},
		points: [][]int{},
	}

		

	//fmt.Println(A.points)


}
