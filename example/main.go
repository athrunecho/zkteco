package main

import (
	"fmt"

	"github.com/athrunecho/zkteco"
	"github.com/northbright/pathhelper"
)

const (
	kaoqingFilePath = "./csv/Attendance/Attendance.csv"
)

func main() {
	var (
		records = [][]string{}
		err     error
		p       string
	)

	if p, err = pathhelper.GetAbsPath(kaoqingFilePath); err != nil {
		fmt.Printf("GetAbsPatherr:%v\n", err)
	}

	if records, err = zkteco.GetCSVRecords(p); err != nil {
		fmt.Printf("records:%v\n", records)
		return
	}
}
