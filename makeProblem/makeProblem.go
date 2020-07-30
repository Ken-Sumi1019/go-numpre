package makeProblem

import (
	"../set"
	"errors"
	"math/rand"
	"time"
)


func init()  {
	rand.Seed(time.Now().UnixNano())
}

func possibleVals(data1 *set.Set,data2 *set.Set,data3 *set.Set) []int {
	data := data1.Product(data2)
	data = data.Product(data3)
	ls := data.ToList()
	result := make([]int,len(ls))
	for i := 0;i < len(ls);i ++ {
		switch val := ls[i].(type) {
		case int:result[i] = val
		}
	}
	return result
}

func shuffle(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func initSet(s *set.Set)  {
	for i := 1;i <= 9;i ++ {
		s.Add(i)
	}
}

func MakeBord() ([][]int,error) {
	side := make([]*set.Set,9)
	vertical := make([]*set.Set,9)
	lattice := make([][]*set.Set,3)
	for i := 0;i < 3;i ++ {
		lattice[i] = make([]*set.Set,3)
		for j := 0;j < 3;j ++ {
			lattice[i][j] = set.MakeSet()
			side[i*3 + j] = set.MakeSet()
			vertical[i*3 + j] = set.MakeSet()
			initSet(lattice[i][j])
			initSet(side[i*3 + j])
			initSet(vertical[i*3+j])
		}
	}
	data := make([][]int,9)
	for i := 0;i < 9;i ++ {
		data[i] = make([]int,9)
	}
	makeBordSolv(data,0,0,side,vertical,lattice)
	return data,nil
}

func makeBordSolv(data [][]int,idx,jdx int,side,vertical []*set.Set,lattice [][]*set.Set) (error) {

	vals := possibleVals(side[idx],vertical[jdx],lattice[idx/3][jdx/3])
	if len(vals) == 0 {
		return errors.New("ダメです")
	} else if len(vals) != 0 && idx == 8 && jdx == 8 {
		data[idx][jdx] = vals[0]
		return nil
	}
	shuffle(vals)
	idx_next := 0;jdx_next := 0
	if jdx == 8 {
		idx_next = idx + 1
		jdx_next = 0
	} else {
		idx_next = idx
		jdx_next = jdx + 1
	}
	for i := 0;i < len(vals);i ++ {
		side[idx].Erase(vals[i]);vertical[jdx].Erase(vals[i]);lattice[idx/3][jdx/3].Erase(vals[i])
		data[idx][jdx] = vals[i]
		woi :=  makeBordSolv(data,idx_next,jdx_next,side,vertical,lattice)
		if woi == nil {
			return nil
		}
		side[idx].Add(vals[i]);vertical[jdx].Add(vals[i]);lattice[idx/3][jdx/3].Add(vals[i])
	}
	return errors.New("ダメです")
}