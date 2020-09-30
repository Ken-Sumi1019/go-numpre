package numpre

import (
	"errors"
	"./set"
	"math/rand"
	"time"
)

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func ProductSetToInt(data1 *set.Set,data2 *set.Set,data3 *set.Set) []int {
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

func initSet(s *set.Set,baseNumber int)  {
	for i := 1;i <= baseNumber*baseNumber;i ++ {
		s.Add(i)
	}
}

func Shuffle(data []int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func Shuffle2D(data [][]int) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

// 完成した盤面を作成する
func MakeBord(baseNumber int) (*Bord,error) {
	side := make([]*set.Set,baseNumber * baseNumber)
	vertical := make([]*set.Set,baseNumber * baseNumber)
	lattice := make([][]*set.Set,baseNumber)
	for i := 0;i < baseNumber;i ++ {
		lattice[i] = make([]*set.Set,baseNumber)
		for j := 0;j < baseNumber;j ++ {
			lattice[i][j] = set.MakeSet()
			side[i*baseNumber + j] = set.MakeSet()
			vertical[i*baseNumber + j] = set.MakeSet()
			initSet(lattice[i][j],baseNumber)
			initSet(side[i*baseNumber + j],baseNumber)
			initSet(vertical[i*baseNumber+j],baseNumber)
		}
	}
	data := make([][]int,baseNumber * baseNumber)
	for i := 0;i < baseNumber * baseNumber;i ++ {
		data[i] = make([]int,baseNumber * baseNumber)
	}
	fillBord(data,0,0,baseNumber,side,vertical,lattice)
	result := &Bord{baseNumber:baseNumber,baseNumberSquare:baseNumber*baseNumber,data:data}
	return result,nil
}

// 再帰的に盤面に数字を埋める
func fillBord(data [][]int,idx,jdx,baseNumber int,side,vertical []*set.Set,lattice [][]*set.Set) (error) {

	vals := ProductSetToInt(side[idx],vertical[jdx],lattice[idx/baseNumber][jdx/baseNumber])
	side_max := baseNumber * baseNumber - 1
	if len(vals) == 0 {
		return errors.New("ダメです")
	} else if len(vals) != 0 && idx == side_max && jdx == side_max {
		data[idx][jdx] = vals[0]
		return nil
	}
	Shuffle(vals)
	idx_next := 0;jdx_next := 0
	if jdx == side_max {
		idx_next = idx + 1
		jdx_next = 0
	} else {
		idx_next = idx
		jdx_next = jdx + 1
	}
	for i := 0;i < len(vals);i ++ {
		side[idx].Erase(vals[i]);vertical[jdx].Erase(vals[i]);lattice[idx/baseNumber][jdx/baseNumber].Erase(vals[i])
		data[idx][jdx] = vals[i]
		woi :=  fillBord(data,idx_next,jdx_next,baseNumber,side,vertical,lattice)
		if woi == nil {
			return nil
		}
		side[idx].Add(vals[i]);vertical[jdx].Add(vals[i]);lattice[idx/baseNumber][jdx/baseNumber].Add(vals[i])
	}
	return errors.New("ダメです")
}