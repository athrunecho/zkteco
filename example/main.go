package main

import (
	"fmt"
	"os"

	"github.com/athrunecho/zkteco"
	"github.com/garyburd/redigo/redis"
	"github.com/northbright/pathhelper"
)

const (
	kaoqinFilePath = "./csv/attendance/attendance.csv"
)

func main() {
	var (
		c   redis.Conn
		arr []string
		k   *zkteco.Kaoqin
		err error
		p   string
		f   *os.File
	)

	if p, err = pathhelper.GetAbsPath(kaoqinFilePath); err != nil {
		fmt.Printf("GetAbsPatherr:%v\n", err)
		return

	}

	if k, err = zkteco.NewKaoqin(":6379", ""); err != nil {
		fmt.Printf("zkteco.NewKaoqin err: %v\n", err)
		return
	}
	if f, err = os.Open(p); err != nil {
		fmt.Printf("os.Open err: %v\n", err)
		return
	}
	defer f.Close()
	if err = k.UpdateAttendances(f); err != nil {
		fmt.Printf("GetCSVRecordserr:%v\n", err)
		return
	}
	defer k.Close()
	c, _ = redis.Dial("tcp", ":6379")
	// key pattern as like "kaoqin:Name"
	if arr, err = redis.Strings(c.Do("HGETALL", "kaoqin:ben")); err != nil {
		fmt.Printf("redis.Strings err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", arr)
}
