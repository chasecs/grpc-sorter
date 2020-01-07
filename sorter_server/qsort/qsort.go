package qsort

func sort(arr []int64, startIndex int, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := startIndex + (startIndex-endIndex)/2
	arr[pivotIndex], arr[endIndex] = arr[endIndex], arr[pivotIndex]
	i, j := startIndex, endIndex-1
	for i < j {
		if arr[i] > arr[endIndex] {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}
	if arr[i] <= arr[endIndex] {
		i++
	}
	arr[i], arr[endIndex] = arr[endIndex], arr[i]

	sort(arr, startIndex, i-1)
	sort(arr, i+1, endIndex)

}

func QuickSort(arr []int64) {
	sort(arr, 0, len(arr)-1)
}
