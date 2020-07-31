package main

import (
	"./makeProblem"
	"fmt"
)


func main()  {
	cost := make([]int,10)
	data := make([][][]int,10)
	for i := 0;i < 10;i ++ {
		cost[i] = 90
	}
	for i := 0;i < 100;i ++ {
		bord,err := makeProblem.MakeBord()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for j := 0;j < 100;j ++ {
			_,prob := makeProblem.MakeProblem(bord)
			count := 0
			for idx := 0;idx < 81;idx ++ {
				k,l := idx/9,idx%9
				if prob[k][l] != 0 {
					count ++
				}
			}
			for m := 0;m < 10;m ++ {
				if cost[m] > count {
					cost[m] = count
					data[m] = prob
					break
				}
			}
		}
	}

	for i := 0;i < 10;i ++ {
		fmt.Println(i," ",cost[i])
		for j := 0;j < 9;j ++ {
			fmt.Println(data[i][j])
		}
	}
}