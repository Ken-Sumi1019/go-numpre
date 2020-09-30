package numpre

import (
	"fmt"
)

type Bord struct {
	baseNumber int
	baseNumberSquare int
	data [][]int
}

func (bord *Bord) OutputBord() [][]int {
	return bord.data
}

func (bord *Bord) StdOutput()  {
	for i,row := range bord.OutputBord() {
		for j,v := range row {
			if 0 <= v && v < 10 {
				fmt.Print(" ")
			}
			fmt.Print(v)
			if j != bord.baseNumberSquare - 1 {
				if (j + 1) % bord.baseNumber == 0 {
					fmt.Print("|")
				} else {
					fmt.Print(",")
				}
			}
		}
		fmt.Println("")
		if (i + 1) % bord.baseNumber == 0 && i != bord.baseNumberSquare - 1 {
			buf := make([]byte,bord.baseNumberSquare*3-1)
			for i := range buf {
				buf[i] = '-'
			}
			fmt.Println(string(buf))
		}
	}
}