package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
// Schedule("7/25/2019 13:45:00")
// Output: 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	const rt = "1/02/2006 15:04:05"
	t, _ := time.Parse(rt, date)
	return t
}

// HasPassed returns whether a date has passed
// HasPassed("July 25, 2019 13:45:00")
// Output: true
func HasPassed(date string) bool {
	const rt = "January 2, 2006 15:04:05"
	at, e := time.Parse(rt, date)
	if e != nil {
		fmt.Println(nil, e)
	}
	ct := time.Now()
	r := at.Before(ct)
	return r
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
// Output: true
func IsAfternoonAppointment(date string) bool {
	const rt = "Monday, January 2, 2006 15:04:05"
	at, e := time.Parse(rt, date)
	if e != nil {
		fmt.Println(nil, e)
	}
	h := at.Hour()
	if h >= 12 && h <= 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
// Description("7/25/2019 13:45:00")
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	const rt = "1/2/2006 15:04:05"
	at, _ := time.Parse(rt, date)
	dw := at.Weekday()
	m := at.Month()
	dn := at.Day()
	y := at.Year()
	h := at.Hour()
	min := at.Minute()
	r := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dw, m, dn, y, h, min)
	return r
}

// AnniversaryDate returns a Time with this year's anniversary
// Output: 2020-09-15
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	// const rd = "2006-1-2"
	// n := time.Now()
	// a := n.AddDate(0, 0, 1)
	// y, m, d := a.Date()
	// r := fmt.Sprintf("%d-%d-%d", y, m, d)
	// result, _ := time.Parse(rd, r)
	// return result

}
