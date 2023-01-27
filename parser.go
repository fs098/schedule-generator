package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// parseLine parses a line of sample-schedule.txt file
func parseLine(line string, getWeekday map[string]Weekday) []Class {
	var (
		sl      = strings.Split(line, ":")
		subject = sl[0]
		agenda  = strings.Split(sl[1], ",")
		result  []Class
	)

	for _, v := range agenda {
		var (
			dayAndTime = strings.Split(v, " ")
			day        = dayAndTime[0]
			time       = dayAndTime[1]
		)

		start, err := strconv.Atoi(time)
		if err != nil {
			log.Fatal(err)
		}

		weekday, ok := getWeekday[day]
		if !ok {
			log.Fatalf("Invalid weekday: %s\n", day)
		}

		class := Class{
			Subject:   subject,
			Day:       weekday,
			StartTime: start,
		}
		result = append(result, class)
	}
	return result
}

// parseClasses parses sample-schedule.txt file
func parseClasses(filename string) [][]Class {
	getWeekday := map[string]Weekday{
		"Mon": Monday,
		"Tue": Tuesday,
		"Wed": Wednesday,
		"Thu": Thursday,
		"Fri": Friday,
		"Sat": Saturday,
		"Sun": Sunday,
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open %s\n", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	classes := [][]Class{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if line[0] == '#' {
			continue
		}
		classes = append(classes, parseLine(line, getWeekday))
	}
	return classes
}
