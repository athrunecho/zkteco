package zkteco

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func GetField(date string) (after string) {
	var (
		year, month, day int64
		err              error
		datePattern      = `^(\d{4}).(\d*).(\d*)$`
	)

	re := regexp.MustCompile(datePattern)
	arr := re.FindStringSubmatch(date)
	if year, err = strconv.ParseInt(arr[1], 10, 32); err != nil {
		return
	}

	if month, err = strconv.ParseInt(arr[2], 10, 32); err != nil {
		return
	}

	if day, err = strconv.ParseInt(arr[3], 10, 32); err != nil {
		return
	}

	t := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.UTC)
	n := t.Unix()
	n = n + (24 * 60 * 60)
	t = time.Unix(n, 0)
	after = fmt.Sprintf("%v\n", t)
	return after
}
