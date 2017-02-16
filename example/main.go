package main

import (
	"fmt"

	"github.com/athrunecho/zkteco"
	"github.com/northbright/pathhelper"
)

const (
	kaoqingFilePath = "./csv/attendance/attendance.csv"
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
		fmt.Printf("GetCSVRecordserr:%v\n", err)
	} else {
		fmt.Printf("%v\n", records)
	}
}
