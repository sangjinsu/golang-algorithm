package level1

func solution(lottos []int, win_nums []int) []int {
	var zeroCnt, matched int
	rank := []int{6, 6, 5, 4, 3, 2, 1}

	for _, lotto := range lottos {
		if lotto == 0 {
			zeroCnt++
			continue
		}
		for _, num := range win_nums {
			if lotto == num {
				matched++
			}
		}
	}

	return []int{rank[matched+zeroCnt], rank[matched]}
}
