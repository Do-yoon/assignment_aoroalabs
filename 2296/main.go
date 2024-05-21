package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Building struct {
	x int
	y int
	c int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(buildings []Building) int {
	// 각 구역별 이익을 저장할 변수 초기화
	profit_asc := 0
	profit_desc := 0

	// x 좌표 기준으로 정렬하여 최대 이익 계산
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x < buildings[j].x
	})
	profit_asc = calcMaxProfit(buildings)

	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x > buildings[j].x
	})
	profit_desc = calcMaxProfit(buildings)

	// 두 가지 경우 중 최대값 반환
	return max(profit_asc, profit_desc)
}

func calcMaxProfit(buildings []Building) int {
	n := len(buildings)

	// DP 배열 초기화
	dp := make([]int, n)

	// DP를 이용하여 최대 이익 계산
	for i := 0; i < n; i++ {
		dp[i] = buildings[i].c // 현재 건물의 이익을 기본값으로 설정

		for j := 0; j < i; j++ {
			// 현재 건물과 비교하여 조건을 만족하는 경우에 한해 최대 이익 갱신
			if buildings[j].y < buildings[i].y {
				dp[i] = max(dp[i], dp[j]+buildings[i].c)
			}
		}
	}

	// 최대 이익 반환
	maxProfit := 0
	for _, profit := range dp {
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var N int
	fmt.Fscanln(reader, &N)

	buildings := make([]Building, N)
	for i := 0; i < N; i++ {
		var x, y, c int
		fmt.Fscanln(reader, &x, &y, &c)
		buildings[i] = Building{x, y, c}
	}

	// 최대 이익 계산
	result := maxProfit(buildings)
	fmt.Println(result)
}
