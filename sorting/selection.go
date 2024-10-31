package sorting

func (a *Algorithms) SelectionSorting(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {

		nMin := i
		temp := arr[i]

		for j := i + 1; j < length; j++ {
			if arr[j] < arr[nMin] {
				nMin = j
			}
		}

		arr[i] = arr[nMin]
		arr[nMin] = temp
	}

	return arr
}
