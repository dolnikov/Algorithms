package sorting

import "fmt"

func (a *Algorithms) SelectionSorting() {
	length := len(a.dataSet)
	for i := 0; i < length; i++ {

		nMin := i
		temp := a.dataSet[i]

		for j := i + 1; j < length; j++ {
			if a.dataSet[j] < a.dataSet[nMin] {
				nMin = j
			}
		}

		a.dataSet[i] = a.dataSet[nMin]
		a.dataSet[nMin] = temp
	}

	fmt.Println(a.dataSet)
}
