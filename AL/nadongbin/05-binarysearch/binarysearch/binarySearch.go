package binarysearch

func binarySearchRecursion(arr []int, target, start, end int) int {
	if start > end {
		return -1
	}
	var mid int
	mid = (start + end) / 2

	if arr[mid] == target {
		return mid + 1
	} else if arr[mid] < target {
		return binarySearchRecursion(arr, target, mid+1, end)
	} else {
		return binarySearchRecursion(arr, target, start, mid-1)
	}
}

func binarySearchLoop(arr []int, target, start, end int) int {
	for start <= end {
		var mid int
		mid = (start + end) / 2

		if arr[mid] == target {
			return mid + 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
