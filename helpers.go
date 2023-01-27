package main

import (
	"fmt"
	"sort"
)

// WastedHours checks how many hours without class exist in a given schedule, between classes
func WastedHours(s Schedule) int {
	var (
		result        = 0
		timeByDay     = TimeByDay(s)
		classDuration = 2
	)

	for _, daySchedule := range timeByDay {
		sort.Ints(daySchedule)
		for i := 0; i < len(daySchedule)-1; i++ {
			// Check if there are any wasted hours between the current class and the next
			result += daySchedule[i+1] - (daySchedule[i] + classDuration)
		}
	}
	return result
}

// TimeByDay returns all class start times of a schedule grouped by day
func TimeByDay(s Schedule) [][]int {
	result := (make([][]int, 7))
	for _, c := range s {
		result[c.Day] = append(result[c.Day], c.StartTime)
	}
	return result
}

// CountHour counts the number of classes that start at the given hour
func CountHour(s Schedule, hour int) int {
	result := 0
	for _, c := range s {
		if c.StartTime == hour {
			result++
		}
	}
	return result
}

// SortByDate sorts the schedules by day and startTime
func SortByDate(s *Schedule) {
	sort.Slice((*s), func(i, j int) bool {
		if (*s)[i].Day != (*s)[j].Day {
			return (*s)[i].Day < (*s)[j].Day
		}
		return (*s)[i].StartTime < (*s)[j].StartTime
	})
}

// PrintN prints the n first schedules, sorted by date
func PrintN(schedules []Schedule, n int) {
	for i := 0; i < n; i++ {
		s := schedules[i]
		SortByDate(&s)
		PrintSchedule(s)
		fmt.Printf("\n")
	}
}

// PrintSchedule prints a Schedule to console
func PrintSchedule(s Schedule) {
	for _, c := range s {
		PrintClass(c)
	}
}

// PrintClass prints a class to console
func PrintClass(c Class) {
	fmt.Printf("%v: %v %v\n", c.Subject, c.Day, c.StartTime)
}
