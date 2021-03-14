package binarysearch

import "testing"

func TestBinarySearchRecursion(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			target: 7,
			want:   4,
		},
		{
			arr:    []int{1, 3, 5, 8, 9, 11, 13, 15, 17, 19},
			target: 7,
			want:   -1,
		},
	}
	for _, tC := range testCases {
		t.Run("이진탐색 재귀", func(t *testing.T) {
			got := binarySearchRecursion(tC.arr, tC.target, 0, len(tC.arr)-1)
			if got != tC.want {
				t.Errorf("Got %d but answer was %d", got, tC.want)
			}
		})
	}
}

func TestBinarySearchLoop(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			arr:    []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
			target: 7,
			want:   4,
		},
		{
			arr:    []int{1, 3, 5, 8, 9, 11, 13, 15, 17, 19},
			target: 7,
			want:   -1,
		},
	}
	for _, tC := range testCases {
		t.Run("이진탐색 재귀", func(t *testing.T) {
			got := binarySearchLoop(tC.arr, tC.target, 0, len(tC.arr)-1)
			if got != tC.want {
				t.Errorf("Got %d but answer was %d", got, tC.want)
			}
		})
	}
}
