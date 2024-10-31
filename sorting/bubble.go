package sorting

import "fmt"

// BubbleSorting
// Time Complexity = O(N^2)
// Space Complexity = O(1)
func (a *Algorithms) BubbleSorting() {

	length := len(a.dataSet)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j+1 < length && a.dataSet[j] > a.dataSet[j+1] {
				temp := a.dataSet[j]
				a.dataSet[j] = a.dataSet[j+1]
				a.dataSet[j+1] = temp
			}
		}
	}

	fmt.Println(a.dataSet)
}
