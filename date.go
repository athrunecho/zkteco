package zkteco

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// GetStartDate returns start time of attendance by given a time range.
func GetStartDate(timeRange string) (t time.Time) {
	var (
		loc              *time.Location
		year, month, day int64
		err              error
		datePattern      = `^(\d{4}).(\d*).(\d*)`
	)

	re := regexp.MustCompile(datePattern)
	arr := re.FindStringSubmatch(timeRange)
	if year, err = strconv.ParseInt(arr[1], 10, 32); err != nil {
		return
	}

	if month, err = strconv.ParseInt(arr[2], 10, 32); err != nil {
		return
	}

	if day, err = strconv.ParseInt(arr[3], 10, 32); err != nil {
		return
	}
	loc, err = time.LoadLocation("Local")

	t = time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, loc)
	return t

}

// GetField adds days and returns according to i.
func GetField(t time.Time, i int) (field string) {
	n := t.AddDate(0, 0, i)
	a, b, c := n.Date()
	field = fmt.Sprintf("%v-%v-%v", a, int(b), c)
	return field
}
