package level1

func solution(a []int, b []int) int {
	result := 0
	arrLen := len(a)
	for i := 0; i < arrLen; i++ {
		result += a[i] * b[i]
	}
	return result
}
