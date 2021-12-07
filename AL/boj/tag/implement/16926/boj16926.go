package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M, R int
	fmt.Fscanf(reader, "%d %d %d\n", &N, &M, &R)

	mat := make([][]string, N)
	for i := 0; i < N; i++ {
		line := make([]string, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(reader, &line[j])
		}
		mat[i] = line
	}

	for i := 0; i < R; i++ {
		rotate(N, M, mat)
	}

	for _, line := range mat {
		l := strings.Join(line, " ")
		fmt.Fprintln(writer, l)
	}

}

func rotate(n int, m int, mat [][]string) {
	check := int(math.Min(float64(n), float64(m))) / 2
	for cnt := 0; cnt < check; cnt++ {
		maxN := n - cnt - 1
		maxM := m - cnt - 1

		temp := mat[cnt][cnt]

		for i := cnt; i < maxM; i++ {
			mat[cnt][i] = mat[cnt][i+1]
		}

		for i := cnt; i < maxN; i++ {
			mat[i][maxM] = mat[i+1][maxM]
		}

		for i := maxM; i > cnt; i-- {
			mat[maxN][i] = mat[maxN][i-1]
		}

		for i := maxN; i > cnt; i-- {
			mat[i][cnt] = mat[i-1][cnt]
		}

		mat[cnt+1][cnt] = temp
	}
}
