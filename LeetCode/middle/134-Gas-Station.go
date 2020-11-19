package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	cur_tank, total_tank, start_station := 0, 0, 0
	for i, n := 0, len(gas); i < n; i++ {
		total_tank += gas[i] - cost[i]
		cur_tank += gas[i] - cost[i]
		if cur_tank < 0 {
			cur_tank = 0
			start_station = i + 1
		}
	}
	if total_tank < 0 {
		return -1
	}
	return start_station
}

func main() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}
