package numpre

import (
	//"../support"
	"./set"
	"fmt"
)


func possibleVals(a,b,c *set.Set) []int {
	baseSet := set.MakeSet()
	for i := 1;i <= 9;i ++ {
		baseSet.Add(i)
	}
	result := baseSet.Difference(a).Difference(b).Difference(c)
	final := make([]int,len(result.Data))
	for i,v := range result.ToList() {
		switch c := v.(type) {
		case int:
			final[i] = c
		}
	}
	return final
}

func Uniqueness(data [][]int) int {
	side := make([]*set.Set,9)
	vertical := make([]*set.Set,9)
	lattice := make([][]*set.Set,3)
	for i := 0;i < 3;i ++ {
		lattice[i] = make([]*set.Set,3)
		for j := 0;j < 3;j ++ {
			lattice[i][j] = set.MakeSet()
			side[i*3 + j] = set.MakeSet()
			vertical[i*3 + j] = set.MakeSet()
			for k := 0;k < 9;k ++ {
				if data[i*3+j][k] != 0 {
					side[i*3 + j].Add(data[i*3+j][k])
				}
				if data[k][i*3+j] != 0 {
					vertical[i*3+j].Add(data[k][i*3+j])
				}
				if data[i*3+k/3][j*3+k%3] != 0 {
					lattice[i][j].Add(data[i*3+k/3][j*3+k%3])
				}
			}
		}
	}
	return makeBordSolv(data,0,0,side,vertical,lattice,0)
}


func makeBordSolv(data [][]int,idx,jdx int,side,vertical []*set.Set,lattice [][]*set.Set,ncount int) (int) {

	vals := possibleVals(side[idx],vertical[jdx],lattice[idx/3][jdx/3])

	idx_next := 0;jdx_next := 0
	if jdx == 8 {
		idx_next = idx + 1
		jdx_next = 0
	} else {
		idx_next = idx
		jdx_next = jdx + 1
	}

	if len(vals) == 0 && data[idx][jdx] == 0 {
		data[idx][jdx] = 0
		return ncount
	} else if len(vals) != 0 && idx == 8 && jdx == 8 && data[idx][jdx] == 0 {
		data[idx][jdx] = vals[0]
		return ncount+1
	} else if idx == 8 && jdx == 8 {
		return ncount + 1
	} else if data[idx][jdx] != 0 {
		return makeBordSolv(data,idx_next,jdx_next,side,vertical,lattice,ncount)
	}

	for i := 0;i < len(vals);i ++ {
		side[idx].Add(vals[i]);vertical[jdx].Add(vals[i]);lattice[idx/3][jdx/3].Add(vals[i])
		data[idx][jdx] = vals[i]
		ncount =  makeBordSolv(data,idx_next,jdx_next,side,vertical,lattice,ncount)
		side[idx].Erase(vals[i]);vertical[jdx].Erase(vals[i]);lattice[idx/3][jdx/3].Erase(vals[i])
	}
	data[idx][jdx] = 0
	return ncount
}

func BordMake() ([][]int) {
	indexes := make([]int,81)
	for i := 0;i < 81;i ++ {
		indexes[i] = i
	}
	Shuffle(indexes)

	bord,err := MakeBord()

	if err != nil {
		fmt.Println(err)
	}

	for i := 0;i < 81;i ++ {
		idx,jdx := indexes[i]/9,indexes[i]%9

		ref := bord[idx][jdx]

		bord[idx][jdx] = 0

		n := Uniqueness(bord)
		if n != 1 {
			bord[idx][jdx] = ref
		}
	}

	return bord
}