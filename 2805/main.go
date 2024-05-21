package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var m int

	fmt.Fscanln(reader, &n, &m)

	// 둘째 줄 입력
	h_tmp, _ := reader.ReadString('\n')
	h_tmp = strings.TrimSpace(h_tmp)
	h_strs := strings.Split(h_tmp, " ")
	h_tree := make([]int, n)
	for i := 0; i < n; i++ {
		h_tree[i], _ = strconv.Atoi(h_strs[i])
	}

	answer := solution(m, h_tree)
	fmt.Fprintln(writer, answer)
}

func solution(m int, h_tree []int) int {
	// 이분 탐색 초기값 설정
	start := 0
	end := max(h_tree)
	answer := 0

	for start <= end {
		mid := (start + end) / 2
		total := 0

		// 현재 높이(mid)에서 잘라낼 수 있는 나무 길이 계산
		for _, height := range h_tree {
			if height > mid {
				total += height - mid
			}
		}

		// 필요한 나무 길이와 비교
		if total >= m {
			answer = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return answer
}

// 배열의 최댓값을 구하는 함수
func max(arr []int) int {
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
