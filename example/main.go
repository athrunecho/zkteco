package main

import (
	"fmt"
	"os"

	"github.com/athrunecho/zkteco"
	"github.com/northbright/pathhelper"
)

const (
	kaoqinFilePath = "./csv/attendance/attendance.csv"
)

func main() {
	var (
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

}
