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

	fmt.Println("Printing Sample:")
	PrintN(schedules, 5)
}
