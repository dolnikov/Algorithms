package sorting

type Algorithms struct {
	dataSet []int
}

func NewSorting() *Algorithms {
	return &Algorithms{
		dataSet: []int{1, 5, 10, 11, 31, 99, 2, 55, 42},
	}
}
