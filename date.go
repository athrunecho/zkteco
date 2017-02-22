package zkteco

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

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
	fmt.Printf("GetStartDate timeRange: %v, t:%v\n", timeRange, t)
	return t

}

func GetField(t time.Time, i int) (field string) {
	fmt.Printf("GetField t:%v, i:%v\n", t, i)
	n := t.Unix()

	n = n + int64(i)*(24*60*60)
	t = time.Unix(n, 0)
	fmt.Printf("t: %v\n", t)
	a, b, c := t.Date()
	field = fmt.Sprintf("%v-%v-%v", a, int(b), c)
	return field
}
