package dsa

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
			swap(arr, j, partition)
		}
	}
	partition += 1
	swap(arr, partition, hi)
	return partition
}

func swap(arr []int, j, i int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}
