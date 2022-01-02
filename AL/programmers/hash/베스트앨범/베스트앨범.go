package 베스트앨범

import (
	"sort"
)

func solution(genres []string, plays []int) []int {

	var answer []int

	genreTotalScore := make(map[string]int)
	songList := make(map[string][]struct {
		idx  int
		play int
	})
	for i := 0; i < len(genres); i++ {
		genreTotalScore[genres[i]] += plays[i]
		songList[genres[i]] = append(songList[genres[i]], struct {
			idx  int
			play int
		}{idx: i, play: plays[i]})
	}

	var sortedTotalScore []struct {
		genre string
		play  int
	}

	for genre, score := range genreTotalScore {
		sortedTotalScore = append(sortedTotalScore, struct {
			genre string
			play  int
		}{genre: genre, play: score})
	}

	sort.Slice(sortedTotalScore, func(i, j int) bool {
		return sortedTotalScore[i].play > sortedTotalScore[j].play
	})

	for _, album := range sortedTotalScore {
		var sortedSongList []struct {
			idx  int
			play int
		}

		for _, song := range songList[album.genre] {
			sortedSongList = append(sortedSongList, song)
		}

		sort.Slice(sortedSongList, func(i, j int) bool {
			return sortedSongList[i].play > sortedSongList[j].play
		})

		if len(sortedSongList) > 1 {
			for _, s := range sortedSongList[:2] {
				answer = append(answer, s.idx)
			}
		} else if len(sortedSongList) == 1 {
			answer = append(answer, sortedSongList[0].idx)
		}

	}

	return answer
}
