package quicksort

func quicksort(arr []int) {
	quicksortHelper(arr, 0, len(arr)-1)
}

func quicksortHelper(arr []int, lo, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		quicksortHelper(arr, lo, p-1)
		quicksortHelper(arr, p+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	partition := lo - 1
	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			partition += 1
			arr[partition], arr[j] = arr[j], arr[partition]
		}
	}
	partition += 1
	arr[partition], arr[hi] = arr[hi], arr[partition]
	return partition
}
