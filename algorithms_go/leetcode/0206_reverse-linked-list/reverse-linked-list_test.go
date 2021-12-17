package leetcode

import (
	"algorithms_go/structures"
	"fmt"
	"testing"
)

type question206 struct {
	para206
	ans206
}

//参数
type para206 struct {
	one []int
}

//答案
type ans206 struct {
	one []int
}

func Test_Problem206(t *testing.T) {
	qs := []question206{
		{
			para206{[]int{5, 4, 3, 2, 1}},
			ans206{[]int{1, 2, 3, 4, 5}},
		},
	}
	fmt.Printf("---------------------Leetcode Problem 206------------------------\n")
	for _, q := range qs {
		_, p := q.ans206, q.para206
		fmt.Printf("【intput】:%v       ", p)
		fmt.Printf("【output】:%v\n", structures.Lists2Ints(reverseList(structures.Ints2Lists(p.one))))
	}
	fmt.Printf("\n\n\n")
}
