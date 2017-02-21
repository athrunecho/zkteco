package zkteco

import (
	"io"
	"regexp"
	//"strconv"

	"github.com/garyburd/redigo/redis"
)

type Kaoqin struct {
	redisAddr     string
	redisPassword string
	c             redis.Conn
}

type AbnormalRecord struct {
	Date               string
	ExceptedsBeginTime string
	ExceptedEndTime    string
	ActualBeginTime    string
	ActualEndTime      string
	Category           string
}

// Update Attendances Records To Redis field := fmt.Printf("%v-%02d", re.FindStringSubmatch(date[2]), k)
func (k *Kaoqin) UpdateAttendances(r io.Reader) (err error) {
	var (
		d           string
		timePattern = `\d{2}.\d{4}.\d{2}|\d{2}.\d{2}`
		records     = [][]string{}
	)

	if records, err = GetCSVRecords(r); err != nil {
		debugPrintf("GetCSVRecordserr:%v\n", err)
		return
	}
	//key
	l := len(records)
	for i := 4; i <= l-1; i += 2 {
		column := records[i]
		maxcolumn := records[0]
		for j := 1; j <= len(maxcolumn); j++ {
			d = GetField(records[2][2])
			row := records[i+1]
			if row[j] != "" {
				re := regexp.MustCompile(timePattern)
				v := re.FindString(row[j])
				k.c.Do("HSET", column[10], d, v)
			} else {
				continue
			}
		}
	}
	return nil
}

func (k *Kaoqin) UpdateArrangements(records [][]string) (err error) {

	//k.c.Do(HSET key field value)
	return nil
}

func (k *Kaoqin) GetAbnormalRecords(name, beginDate, endDate string) (records []AbnormalRecord, err error) {
	return []AbnormalRecord{}, nil

}
