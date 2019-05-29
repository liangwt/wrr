package main

import (
	"fmt"
	"github.com/liangwt/wrr"
)

func main() {
	points := []*wrr.Point{
		{Entry: "A", Weight: 5},
		{Entry: "B", Weight: 2},
		{Entry: "C", Weight: 3},
	}

	iter := wrr.NewWrr(points)
	for i := 0; i < 12; i++ {
		fmt.Printf("%s ", iter.Next().Entry)
	}

	fmt.Println()

	smIter := wrr.NewSmoothWrr(points)
	for i := 0; i < 12; i++ {
		fmt.Printf("%s ", smIter.Next().Entry)
	}
}
