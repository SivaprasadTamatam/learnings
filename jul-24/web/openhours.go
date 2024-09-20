package main

import (
	"fmt"
	"sort"
)

// daysOfWeek defines the days of the week in order
var daysOfWeek = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// GroupOpenHours transforms the open hours data to the grouped format
func GroupOpenHours(openHours []map[string]string) []map[string]string {
	// Create a lookup map for day indices
	dayIndex := map[string]int{}
	for i, day := range daysOfWeek {
		dayIndex[day] = i
	}

	// Sort the input data based on the day of the week
	sort.Slice(openHours, func(i, j int) bool {
		return dayIndex[openHours[i]["day"]] < dayIndex[openHours[j]["day"]]
	})

	// Create a map to hold the open hours for each day
	openHoursMap := make(map[string]map[string]string)
	for _, entry := range openHours {
		openHoursMap[entry["day"]] = map[string]string{
			"open":  entry["open"],
			"close": entry["close"],
		}
	}

	// Group consecutive days with the same open and close times
	grouped := []map[string]string{}
	start := 0
	for start < len(daysOfWeek) {
		day := daysOfWeek[start]
		openTime := ""
		closeTime := ""

		if hours, exists := openHoursMap[day]; exists {
			openTime = hours["open"]
			closeTime = hours["close"]
		}

		end := start
		for end+1 < len(daysOfWeek) {
			nextDay := daysOfWeek[end+1]
			if nextHours, exists := openHoursMap[nextDay]; exists {
				if nextHours["open"] != openTime || nextHours["close"] != closeTime {
					break
				}
			} else if openTime != "" || closeTime != "" {
				break
			}
			end++
		}

		if start == end {
			grouped = append(grouped, map[string]string{
				"days":  day,
				"open":  openTime,
				"close": closeTime,
			})
		} else {
			grouped = append(grouped, map[string]string{
				"days":  fmt.Sprintf("%s-%s", daysOfWeek[start], daysOfWeek[end]),
				"open":  openTime,
				"close": closeTime,
			})
		}
		start = end + 1
	}
	// Sort the grouped result by the day of the week to ensure order
	sort.Slice(grouped, func(i, j int) bool {
		// Split the day ranges for comparison
		dayRangeI := grouped[i]["days"]
		dayRangeJ := grouped[j]["days"]

		// Get the start day for each range
		startDayI := dayRangeI
		startDayJ := dayRangeJ

		if dashIndex := findDashIndex(dayRangeI); dashIndex != -1 {
			startDayI = dayRangeI[:dashIndex]
		}
		if dashIndex := findDashIndex(dayRangeJ); dashIndex != -1 {
			startDayJ = dayRangeJ[:dashIndex]
		}

		return dayIndex[startDayI] < dayIndex[startDayJ]
	})

	return grouped
}

func main() {
	openHours := []map[string]string{
		{"day": "Monday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Wednesday", "open": "8:00 AM", "close": "6:00 PM"},
		{"day": "Thursday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Friday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Saturday", "open": "8:00 AM", "close": "4:00 PM"},
	}
	expected := []map[string]string{
		{"days": "Monday-Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
		{"days": "Wednesday", "open": "8:00 AM", "close": "6:00 PM"},
		{"days": "Thursday-Friday", "open": "8:00 AM", "close": "5:00 PM"},
		{"days": "Saturday", "open": "8:00 AM", "close": "4:00 PM"},
		{"days": "Sunday", "open": "", "close": ""},
	}
	actual := GroupOpenHours(openHours)
	fmt.Println(actual)
}
