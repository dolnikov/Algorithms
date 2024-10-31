package sorting

import "fmt"

type Algorithms struct {
	dataSet []int
}

func NewSorting() *Algorithms {
	return &Algorithms{
		dataSet: []int{1, 5, 10, 11, 31, 99, 2, 55, 42},
	}
}

func (a *Algorithms) Run() {
	//res := a.BubbleSorting(a.dataSet)
	//res := a.SelectionSorting(a.dataSet)
	res := a.InsertionSorting(a.dataSet)

	fmt.Println(res)
}
