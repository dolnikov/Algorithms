package sorting

func (a *Algorithms) BubbleSorting(arr []int) []int {

	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if j+1 < length && arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}
