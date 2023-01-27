package main

import (
	"reflect"
	"testing"
)

func TestAllSchedules1(t *testing.T) {
	classes := [][]Class{}
	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) != 1 || len(schedules[0]) != 0 {
		t.Error("Expected [[]] but got", schedules)
	}
}

func TestAllSchedules2(t *testing.T) {
	classes := [][]Class{
		{
			{Subject: "pf", Day: Monday, StartTime: 12},
			{Subject: "pf", Day: Tuesday, StartTime: 13},
		},
		{
			{Subject: "tmd", Day: Friday, StartTime: 9},
			{Subject: "tmd", Day: Wednesday, StartTime: 11},
		},
	}
	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) != 4 {
		t.Error("Expected 4 schedules but got", len(schedules))
	}
}

func TestAllSchedules3(t *testing.T) {
	classes := [][]Class{
		{
			{Subject: "pf", Day: Monday, StartTime: 12},
			{Subject: "pf", Day: Tuesday, StartTime: 13},
		},
		{
			{Subject: "tmd", Day: Monday, StartTime: 12},
			{Subject: "tmd", Day: Wednesday, StartTime: 11},
		},
	}
	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) != 3 {
		t.Error("Expected 3 schedules but got", len(schedules))
	}

}

func TestAllSchedules4(t *testing.T) {
	classes := [][]Class{
		{
			{Subject: "pf", Day: Monday, StartTime: 12},
			{Subject: "pf", Day: Tuesday, StartTime: 13},
		},
		{
			{Subject: "tmd", Day: Monday, StartTime: 12},
		},
	}
	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) != 1 {
		t.Error("Expected 1 schedule but got", len(schedules))
	}
}

func TestAllSchedules5(t *testing.T) {
	classes := [][]Class{
		{
			{Subject: "pf", Day: Monday, StartTime: 12},
			{Subject: "pf", Day: Tuesday, StartTime: 13},
		},
		{
			{Subject: "tmd", Day: Friday, StartTime: 9},
			{Subject: "tmd", Day: Wednesday, StartTime: 11},
		},
		{
			{Subject: "ds", Day: Monday, StartTime: 9},
			{Subject: "ds", Day: Friday, StartTime: 14},
		},
	}

	expectedSchedules := []Schedule{
		{{Subject: "pf", Day: Monday, StartTime: 12}, {Subject: "tmd", Day: Friday, StartTime: 9}, {Subject: "ds", Day: Monday, StartTime: 9}},
		{{Subject: "pf", Day: Monday, StartTime: 12}, {Subject: "tmd", Day: Friday, StartTime: 9}, {Subject: "ds", Day: Friday, StartTime: 14}},
		{{Subject: "pf", Day: Monday, StartTime: 12}, {Subject: "tmd", Day: Wednesday, StartTime: 11}, {Subject: "ds", Day: Monday, StartTime: 9}},
		{{Subject: "pf", Day: Monday, StartTime: 12}, {Subject: "tmd", Day: Wednesday, StartTime: 11}, {Subject: "ds", Day: Friday, StartTime: 14}},
		{{Subject: "pf", Day: Tuesday, StartTime: 13}, {Subject: "tmd", Day: Friday, StartTime: 9}, {Subject: "ds", Day: Monday, StartTime: 9}},
		{{Subject: "pf", Day: Tuesday, StartTime: 13}, {Subject: "tmd", Day: Friday, StartTime: 9}, {Subject: "ds", Day: Friday, StartTime: 14}},
		{{Subject: "pf", Day: Tuesday, StartTime: 13}, {Subject: "tmd", Day: Wednesday, StartTime: 11}, {Subject: "ds", Day: Monday, StartTime: 9}},
		{{Subject: "pf", Day: Tuesday, StartTime: 13}, {Subject: "tmd", Day: Wednesday, StartTime: 11}, {Subject: "ds", Day: Friday, StartTime: 14}},
	}

	schedules := AllSchedules(classes, Schedule{}, 0)
	if len(schedules) != len(expectedSchedules) {
		t.Errorf("Expected %d schedules, got %d", len(expectedSchedules), len(schedules))
	}

	for i := 0; i < 8; i++ {
		if !reflect.DeepEqual(schedules[i], expectedSchedules[i]) {
			t.Errorf("Expected %v schedule, got %v", expectedSchedules[i], schedules[i])
		}
	}
}

func TestOverlap(t *testing.T) {
	// Test case 1: Schedule contains a class that overlaps with the given class
	s := Schedule{
		{Subject: "math", Day: Monday, StartTime: 9},
		{Subject: "science", Day: Tuesday, StartTime: 10},
		{Subject: "history", Day: Wednesday, StartTime: 11},
	}
	c := Class{Subject: "english", Day: Wednesday, StartTime: 11}
	if !overlap(s, c) {
		t.Error("Expected overlap")
	}

	// Test case 2: Schedule contains a class that doesn't overlap with the given class
	s = Schedule{
		{Subject: "math", Day: Monday, StartTime: 9},
		{Subject: "science", Day: Tuesday, StartTime: 10},
		{Subject: "history", Day: Wednesday, StartTime: 11},
	}
	c = Class{Subject: "english", Day: Wednesday, StartTime: 12}
	if overlap(s, c) {
		t.Error("Did not expected overlap")
	}

	// Test case 3: Schedule is empty, so no overlap
	s = Schedule{}
	c = Class{Subject: "english", Day: Wednesday, StartTime: 12}
	if overlap(s, c) {
		t.Error("Expected no overlap")
	}
}
