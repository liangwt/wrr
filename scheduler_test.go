package wrr

import "testing"

var maxMatchNum = 500

var groupOne = []*Point{
	{Entry: "A", Weight: 5},
	{Entry: "B", Weight: 2},
	{Entry: "C", Weight: 3},
}

var groupOneResult = []string{"A", "A", "A", "C", "A", "B", "C", "A", "B", "C"}
var groupOneSmoothResult = []string{"A", "C", "B", "A", "A", "C", "A", "B", "C", "A"}

var groupTwo = []*Point{
	{Entry: "A", Weight: 5},
	{Entry: "B", Weight: 0},
	{Entry: "C", Weight: 3},
}

var groupTwoResult = []string{"A", "A", "A", "C", "A", "C", "A", "C"}
var groupTwoSmoothResult = []string{"A", "C", "A", "A", "C", "A", "C", "A"}

var groupThree = []*Point{
	{Entry: "A", Weight: 50},
	{Entry: "B", Weight: 2},
	{Entry: "C", Weight: 3},
	{Entry: "D", Weight: 3},
	{Entry: "E", Weight: 3},
	{Entry: "F", Weight: 3},
	{Entry: "G", Weight: 10},
	{Entry: "H", Weight: 3},
	{Entry: "I", Weight: 3},
}

var groupThreeResult = []string{"A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "A", "G", "A", "G", "A", "G", "A", "G", "A", "G", "A", "G", "A", "G", "A", "C", "D", "E", "F", "G", "H", "I", "A", "B", "C", "D", "E", "F", "G", "H", "I", "A", "B", "C", "D", "E", "F", "G", "H", "I"}
var groupThreeSmoothResult = []string{"A", "A", "G", "A", "C", "A", "A", "D", "A", "E", "A", "A", "G", "A", "F", "A", "H", "A", "A", "I", "A", "A", "G", "A", "A", "B", "A", "A", "G", "A", "A", "C", "A", "D", "A", "A", "G", "A", "E", "A", "A", "F", "A", "H", "A", "A", "G", "A", "I", "A", "A", "A", "G", "A", "A", "B", "A", "A", "G", "A", "C", "A", "A", "D", "A", "E", "A", "F", "A", "A", "G", "A", "H", "A", "A", "I", "A", "A", "G", "A"}

func TestNewWrr(t *testing.T) {
	iter1 := NewWrr(groupOne)

	for i := 0; i < maxMatchNum; i++ {
		m := groupOneResult[i%len(groupOneResult)]
		if iter1.Next().Entry.(string) != m {
			t.Fatalf("group one not match at %d => %s", i, m)
		}
	}

	iter2 := NewWrr(groupTwo)

	for i := 0; i < maxMatchNum; i++ {
		m := groupTwoResult[i%len(groupTwoResult)]
		if iter2.Next().Entry.(string) != m {
			t.Fatalf("group two not match at %d => %s", i, m)
		}
	}

	iter3 := NewWrr(groupThree)

	for i := 0; i < maxMatchNum; i++ {
		m := groupThreeResult[i%len(groupThreeResult)]
		if iter3.Next().Entry.(string) != m {
			t.Fatalf("group two not match at %d => %s", i, m)
		}
	}
}

func TestNewSmoothWrr(t *testing.T) {
	iter1 := NewSmoothWrr(groupOne)

	for i := 0; i < maxMatchNum; i++ {
		m := groupOneSmoothResult[i%len(groupOneSmoothResult)]
		if iter1.Next().Entry.(string) != m {
			t.Fatalf("group one not match at %d => %s", i, m)
		}
	}

	iter2 := NewSmoothWrr(groupTwo)

	for i := 0; i < maxMatchNum; i++ {
		m := groupTwoSmoothResult[i%len(groupTwoSmoothResult)]
		if iter2.Next().Entry.(string) != m {
			t.Fatalf("group two not match at %d => %s", i, m)
		}
	}

	iter3 := NewSmoothWrr(groupThree)

	for i := 0; i < maxMatchNum; i++ {
		m := groupThreeSmoothResult[i%len(groupThreeSmoothResult)]
		if iter3.Next().Entry.(string) != m {
			t.Fatalf("group two not match at %d => %s", i, m)
		}
	}
}
