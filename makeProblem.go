package numpre

import (
	"./set"
	//"../support"
)

func check2DZero(ls [][]int) bool {
	for i := 0;i < len(ls);i ++ {
		for j := 0;j < len(ls[0]);j ++ {
			if ls[i][j] == 0{return true}
		}
	}
	return false
}

// それぞれのマスに入る可能税のある数字を列挙
func possibilityData(confirmBord [][]int,possibleNumbers [][][]int)  {
	side := make([]*set.Set,9)
	vertical := make([]*set.Set,9)
	lattice := make([][]*set.Set,3)
	for i := 0;i < 3;i ++ {
		lattice[i] = make([]*set.Set,3)
		for j := 0;j < 3;j ++ {
			lattice[i][j] = set.MakeSet()
			side[i*3 + j] = set.MakeSet()
			vertical[i*3 + j] = set.MakeSet()
		}
	}
	for i := 0;i < 9;i ++ {
		for j := 0;j < 9;j ++ {
			for k := 0;k < 9;k ++ {
				possibleNumbers[i][j][k] = 0
			}
			if confirmBord[i][j] != 0 {
				side[i].Add(confirmBord[i][j])
				vertical[j].Add(confirmBord[i][j])
				lattice[i/3][j/3].Add(confirmBord[i][j])
			}
		}
	}
	for i := 0;i < 9;i ++ {
		for j := 0;j < 9;j ++ {
			if confirmBord[i][j] != 0 {
				possibleNumbers[i][j][confirmBord[i][j]-1] = 1
				continue
			}
			for k := 0;k < 9;k ++ {
				if (!side[i].Exist(k+1)) && (!vertical[j].Exist(k+1)) && (!lattice[i/3][j/3].Exist(k+1)) {
					possibleNumbers[i][j][k] = 1
				}
			}
		}
	}
}

// 数字を抜いてみて答えが一意に決まる場合には抜く
func (bord_ *Bord) MakeProblem() ([][]int,[][]int) {
	bord := bord_.data
	possibleNumbers := make([][][]int,9)
	confirmBord := make([][]int,9)
	initialBord := make([][]int,9)
	for i := 0;i < 9;i ++ {
		possibleNumbers[i] = make([][]int,9)
		confirmBord[i] = make([]int,9)
		initialBord[i] = make([]int,9)
		for j := 0;j < 9;j ++ {
			possibleNumbers[i][j] = make([]int,9)
		}
	}
	idxList := make([][]int,81)
	for i := 0;i < 9;i ++ {
		for j := 0;j < 9;j ++ {
			idxList[i*9+j] = []int{i,j}
		}
	}
	Shuffle2D(idxList)
	idxidx := 0
	for check2DZero(confirmBord) {
		x,y := 0,0
		for ;idxidx < 81;idxidx++ {
			x,y = idxList[idxidx][0],idxList[idxidx][1]
			if confirmBord[x][y] == 0 {
				break
			}
		}
		confirmBord[x][y] = bord[x][y]
		initialBord[x][y] = bord[x][y]
		check := true
		for check {
			check = false
			possibilityData(confirmBord,possibleNumbers)
			for idx := 0;idx < 81;idx ++ {
				i := idx / 9;j := idx % 9
				if confirmBord[i][j] != 0 {
					continue
				}
				count := 0;numIdx := 0
				for k := 0;k < 9;k ++ {
					if possibleNumbers[i][j][k] == 1 {
						numIdx = k
					}
					count += possibleNumbers[i][j][k]
				}
				if count == 1 {
					confirmBord[i][j] = numIdx + 1
					check = true
					continue
				}
				num := bord[i][j]
				tate,yoko,waku,wakuI,wakuJ := 0,0,0,(i/3)*3,(j/3)*3
				for k := 0;k < 9;k ++ {
					// 縦
					tate += possibleNumbers[k][j][num-1]
					// 横
					yoko += possibleNumbers[i][k][num-1]
					// 枠
					waku += possibleNumbers[wakuI+k/3][wakuJ+k%3][num-1]
				}
				if tate == 1 || yoko == 1 || waku == 1 {
					confirmBord[i][j] = num
					check = true
				}
			}
		}
	}
	return confirmBord,initialBord
}