package numpre

import (
	"fmt"
	"testing"
)

func TestUniqueness(t *testing.T) {
	bord := [][]int{
		{0,8,0,0,9,0,0,0,0},
		{1,4,3,0,0,0,0,0,2},
		{0,7,6,0,3,5,0,0,0},
		{0,2,0,0,0,1,0,5,9},
		{0,0,0,0,0,0,6,0,0},
		{0,0,0,0,0,0,7,0,8},
		{0,0,4,6,8,7,0,0,3},
		{6,0,0,0,0,0,9,0,0},
		{0,0,0,0,4,0,0,6,0},
	}
	num := Uniqueness(bord)
	fmt.Println(num)
}