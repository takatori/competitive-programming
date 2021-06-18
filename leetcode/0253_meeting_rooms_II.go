package main

import "sort"

func minMeetingRooms(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	end_times := make([]int, 0)

	end_times = append(end_times, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		j := 0
		for ; j < len(end_times); j++ {
			if end_times[j] <= interval[0] {
				end_times[j] = interval[1]
				break
			}
		}

		if j == len(end_times) {
			end_times = append(end_times, interval[1])
		}
	}

	return len(end_times)
}
