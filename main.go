package main

import (
	"fmt"
	"os"
)

const sample string = "./sample-schedule.txt"

func main() {
	classes := parseClasses(sample)
	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) == 0 {
		fmt.Println("No schedules were found")
		os.Exit(0)
	}
	fmt.Println("Total schedules found:", len(schedules))

	// filter schedules with minimum wasted hours between classes
	byWastedHours, n := FilterByMin(schedules, WastedHours)
	fmt.Printf("Schedules found: %d, wasted hours: %d\n", len(byWastedHours), n)

	// filter schedules with the least classes at 18h
	min18H, n := FilterByMin(byWastedHours, func(s Schedule) int {
		return CountHour(s, 18)
	})
	fmt.Printf("Schedules found: %d, classes at 18h: %d\n", len(min18H), n)

	// filter schedules with the least classes at 9h
	min9H, count9h := FilterByMin(min18H, func(s Schedule) int {
		return CountHour(s, 9)
	})
	fmt.Printf("Schedules found: %d, classes at 9h: %d, at 18h: %d\n", len(min9H), count9h, n)

	// filter schedules with the least classes at 16h
	min16H, count16h := FilterByMin(min18H, func(s Schedule) int {
		return CountHour(s, 16)
	})
	fmt.Printf("Schedules found: %d, classes at 16h: %d, at 18h: %d\n", len(min16H), count16h, n)

	fmt.Printf("\nPrinting Sample:\n\n")
	PrintN(min9H, 3)
}
