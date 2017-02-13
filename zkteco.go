package zkteco

import (
	"fmt"

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

func (k *Kaoqin) UpdateAttendances(f string) (err error) {

}

func (k *Kaoqin) UpdateArrangements(f string) (err error) {

}

func (k *Kaoqin) GetAbnormalRecords(name, beginDate, endDate string) (records []AbnormalRecord, err error) {

}
