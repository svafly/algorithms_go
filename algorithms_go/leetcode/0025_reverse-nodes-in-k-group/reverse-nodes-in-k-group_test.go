package leetcode

import (
	"algorithms_go/structures"
	"fmt"
	"testing"
)

type question25 struct {
	para25
	ans25
}
type para25 struct {
	one []int
	k   int
}
type ans25 struct {
	one []int
}

func Test_Problem25(t *testing.T) {
	qs := []question25{
		{
			para25{[]int{1, 2, 3, 4, 5}, 2},
			ans25{[]int{2, 1, 4, 3, 5}},
		},
		{
			para25{[]int{1, 2, 3, 4, 5}, 3},
			ans25{[]int{3, 2, 1, 4, 5}},
		},
	}
	fmt.Printf("---------------------Leetcode Problem 25------------------------\n")
	for _, q := range qs {
		_, p := q.ans25, q.para25
		fmt.Printf("【intput】:%v       ", p)
		fmt.Printf("【output】:%v\n", structures.Lists2Ints(reverseKGroup(structures.Ints2Lists(p.one), p.k)))
	}
	fmt.Printf("\n\n\n")
}
