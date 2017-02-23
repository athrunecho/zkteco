package zkteco

import (
	"io"
	"regexp"

	"github.com/garyburd/redigo/redis"
)

// Kaoqin contains Redis address, Redis password and Redis connection
type Kaoqin struct {
	// redisAddr is Redis address. Ex: ":6379", "192.168.1.18:6379".
	redisAddr string
	// redisPassword is Redis password.
	redisPassword string
	// c is Redis connection
	c redis.Conn
}

// NewKaoqin returns an pointer to Kaoqin by given Redis address, Redis password.
func NewKaoqin(redisAddr, redisPassword string) (k *Kaoqin, err error) {
	k = &Kaoqin{redisAddr: redisAddr, redisPassword: redisPassword}
	if k.c, err = GetRedisConn(k.redisAddr, k.redisPassword); err != nil {
		debugPrintf("GetRedisConn err:%v\n", err)
		return &Kaoqin{}, err
	}
	return k, nil
}

// Update attendances records to redis.
func (k *Kaoqin) UpdateAttendances(r io.Reader) (err error) {
	var (
		field       string
		timePattern = `\d{2}.\d{4}.\d{2}|\d{2}.\d{2}`
		records     = [][]string{}
	)

	//  records gets the file that GetCSVRecords returned.
	if records, err = GetCSVRecords(r); err != nil {
		debugPrintf("GetCSVRecords err:%v\n", err)
		return
	}
	l := len(records)
	columnNum := len(records[0])
	t := GetStartDate(records[2][2])

	for i := 4; i <= l-1; i += 2 {
		columns := records[i]
		for j := 0; j <= columnNum-1; j++ {
			field = GetField(t, j)
			row := records[i+1]
			if row[j] != "" {
				re := regexp.MustCompile(timePattern)
				value := re.FindString(row[j])
				if _, err = k.c.Do("HSET", columns[10], field, value); err != nil {
					debugPrintf("k.c.Do err:%v\n", err)
				}
			}
		}
	}
	return
}
