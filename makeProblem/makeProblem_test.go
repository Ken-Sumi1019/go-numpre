package makeProblem

import (
	"fmt"
	"testing"
	"../set"
)

func TestMakeBord(t *testing.T) {
	for i := 0;i < 10;i ++ {
		data,err := MakeBord()
		if err != nil {
			t.Error(err)
			return
		}
		// 整合性チェック
		side := set.MakeSet()
		v := make([]*set.Set,9)
		l := make([][]*set.Set,3)
		for i := 0;i < 3;i ++ {
			l[i] = make([]*set.Set,3)
			for j := 0;j < 3;j ++ {
				l[i][j] = set.MakeSet()
			}
		}
		for i := 0;i < 9;i ++ {
			v[i] = set.MakeSet()
		}
		for i := 0;i < 9;i ++ {
			for j := 0;j < 9;j ++ {
				side.Add(data[i][j])
				v[j].Add(data[i][j])
				l[i/3][j/3].Add(data[i][j])
			}
			if len(side.ToList()) != 9 {
				fmt.Println(data)
				t.Error("side 数字が一意でありません")
				return
			}
		}
		for i := 0;i < 9;i ++ {
			if len(v[i].ToList()) != 9 {
				fmt.Println(data)
				t.Error("v 数字が一意でありません")
				return
			}
			if len(l[i/3][i%3].ToList()) != 9 {
				fmt.Println(data)
				t.Error("l 数字が一意でありません")
				return
			}
		}
	}
}