package main

import (
	"fmt"
	"inventory_series/utilities"

	"os"

	"gonum.org/v1/plot"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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


func generatePoints (arrayOfPoints []int) plotter.XYs {
	pts := make(plotter.XYs, len(arrayOfPoints))
	for i:= range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(arrayOfPoints[i])
	}
	return pts
}


func main () {
	var result [][]int = inventorySteps(500)

	//Writing down the result
	file,err := os.Create("inventory.txt")
	utilities.CheckErr(err)
	for _,v:=range result {
		_,err := file.Write(utilities.StringArray(v))
		utilities.CheckErr(err)
	}
	closeErr := file.Close()
	utilities.CheckErr(closeErr)

	//Plotting the result
	p := plot.New()
	p.Title.Text = "Inventory sequence"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for i,v := range result {
		plotError := plotutil.AddLinePoints(p,fmt.Sprint(i),generatePoints(v),plotutil.Color(50))
		utilities.CheckErr(plotError)
	}
	p.Save(10*vg.Inch,10*vg.Inch,"points.png")



	
}
