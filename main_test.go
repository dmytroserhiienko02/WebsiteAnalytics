package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetRecords(t *testing.T) {
	expected := map[int]map[int]bool{
		1: {101: true},
		4: {104: true},
		5: {105: true, 106: true},
	}

	result := GetRecords("days/day2.csv")

	fmt.Printf("Actual result: %v\n", result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GetRecords() failed, expected: %v, got: %v", expected, result)
	}
}

func TestDoubleVisited(t *testing.T) {
	// Prepare test data
	firstDay := map[int]map[int]bool{
		1: {101: true},
		2: {102: true, 103: true},
	}
	secondDay := map[int]map[int]bool{
		2: {102: true, 103: true},
		3: {101: true},
	}
	expected := []int{2}

	result := DoubleVisited(firstDay, secondDay)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DoubleVisited() failed, expected: %v, got: %v", expected, result)
	}
}

func TestParseRec(t *testing.T) {
	rec := []string{"1", "101", "2024-01-01T00:00:00"}
	expected := Record{
		UserID:    1,
		ProductID: 101,
		Timestamp: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	result, err := parseRec(rec)

	if err != nil {
		t.Errorf("parseRec() failed with error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parseRec() failed, expected: %v, got: %v", expected, result)
	}
}
