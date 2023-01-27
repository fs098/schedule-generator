package main

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Class struct {
	Subject   string
	Day       Weekday
	StartTime int
}

type Schedule []Class

// allSchedules returns all possible schedules that can be formed from a set of classes.
// Each element of classes contains an array with all the time slots of a class, for example:
// classes[0] = [{pf, monday, 12h}, {pf, tuesday, 13h}, {pf...}...],
// classes[1] = [{tmd, friday, 9h}, {tmd, wednesday, 11h}, {tmd...}...]...
// s starts as an empty schedule and index at 0, we iterate through the classes at the current index
// and, when we find one that does not overlap, we then call allSchedules recursively, updating both
// s and index variables. When index == len(classes) we have successfully created a schedule without
// overlapping classes.
func AllSchedules(classes [][]Class, s Schedule, index int) []Schedule {
	if index == len(classes) {
		return []Schedule{s}
	}

	var result []Schedule
	for _, c := range classes[index] {
		if overlap(s, c) {
			continue
		}
		new := make(Schedule, len(s))
		copy(new, s)
		new = append(new, c)
		result = append(result, AllSchedules(classes, new, index+1)...)
	}
	return result
}

// overlap checks if given class overlaps with the schedule, because all classes in
// the sample schedule start at the same time and have the same time slot, only
// the start time needs to be considered
func overlap(s Schedule, c Class) bool {
	for _, sc := range s {
		if sc.Day == c.Day && sc.StartTime == c.StartTime {
			return true
		}
	}
	return false
}
