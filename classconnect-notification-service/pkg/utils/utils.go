package utils

import (
	"time"
)

func IsEvenWeek() bool {
	_, week := time.Now().ISOWeek()
	return week%2 == 0
}

func GetDayOfWeek() int {
	return int(time.Now().Weekday())
}

// TODO: replace crutch
func GetDifferentTimeRFFC822(t string) (int, error) {
	layout := "15:04"
	firstParse, err := time.Parse(layout, t)
	if err != nil {
		return -1, err
	}

	secondParse, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		return -1, err
	}

	return int(firstParse.Sub(secondParse).Minutes()), nil
}
