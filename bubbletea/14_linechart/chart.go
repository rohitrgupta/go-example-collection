package main

import (
	"fmt"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart/streamlinechart"
	"github.com/NimbleMarkets/ntcharts/linechart/timeserieslinechart"
)

func main() {
	slc := streamlinechart.New(13, 10)
	for _, v := range []float64{4, 6, 8, 10, 8, 6, 4, 2, 0, 2, 4} {
		slc.Push(v)
	}
	slc.Draw()

	fmt.Println(slc.View())

	tslc := timeserieslinechart.New(41, 10)
	for i, v := range []float64{0, 4, 8, 10, 8, 4, 0, -4, -8, -10, -8, -4, 0} {
		date := time.Now().Add(time.Hour * time.Duration(24*i))
		tslc.Push(timeserieslinechart.TimePoint{date, v})
	}
	tslc.DrawBraille()

	fmt.Println(tslc.View())
}
