package zkteco

import (
	//"fmt"
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

func NewKaoqin(redisAddr, redisPassword string) (k *Kaoqin, err error) {
	k = &Kaoqin{redisAddr: redisAddr, redisPassword: redisPassword}
	if k.c, err = GetRedisConn(k.redisAddr, k.redisPassword); err != nil {
		debugPrintf("GetRedisConn err:%v\n", err)
		return &Kaoqin{}, err
	}
	return k, nil
}

// Update Attendances Records To Redis
func (k *Kaoqin) UpdateAttendances(r io.Reader) (err error) {
	var (
		d           string
		timePattern = `\d{2}.\d{4}.\d{2}|\d{2}.\d{2}`
		records     = [][]string{}
	)

	if records, err = GetCSVRecords(r); err != nil {
		debugPrintf("GetCSVRecords err:%v\n", err)
		return
	}
	//key
	l := len(records)
	columnNum := len(records[0])
	t := GetStartDate(records[2][2])

	for i := 4; i <= l-1; i += 2 {
		columns := records[i]
		for j := 0; j <= columnNum-1; j++ {
			d = GetField(t, j)
			row := records[i+1]
			if row[j] != "" {
				re := regexp.MustCompile(timePattern)
				v := re.FindString(row[j])
				if _, err = k.c.Do("HSET", columns[10], d, v); err != nil {
					debugPrintf("k.c.Do err:%v\n", err)
				}
			}
		}
	}
	return
}

func (k *Kaoqin) UpdateArrangements(records [][]string) (err error) {

	//k.c.Do(HSET key field value)
	return nil
}

func (k *Kaoqin) GetAbnormalRecords(name, beginDate, endDate string) (records []AbnormalRecord, err error) {
	return []AbnormalRecord{}, nil

}
