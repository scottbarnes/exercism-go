package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = 0 // Not used. :(
)

// Returns the correct day of the month, regardless of the WeekSchedule (save for Last)
func getDay(week WeekSchedule, year int, month time.Month, weekday time.Weekday) int {
	for day := int(week); ; day++ {
		t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if t.Weekday() == weekday {
			return t.Day()
		}
	}
}

// Get the date of the last 'weekday' in the month.
func getLast(year int, month time.Month, weekday time.Weekday) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	t = t.AddDate(0, 1, -7)
	for t.Weekday() != weekday {
		t = t.AddDate(0, 0, 1)
	}
	return t.Day()
}

// Return the relevant day, as an int, of weekday.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	switch week {
	case Last:
		return getLast(year, month, weekday)
	default:
		// Everything but last.
		return getDay(week, year, month, weekday)
	}
}
