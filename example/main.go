package main

import (
	"fmt"

	"github.com/athrunecho/zkteco"
	"github.com/northbright/pathhelper"
)

var (
	err error
	p   string
)

const (
	kaoqingFilePath = "./in/kaoqing.csv"
)

func main() {
	var records = [][]string{}

	if p, err = pathhelper.GetAbsPath(kaoqingFilePath); err != nil {
		fmt.Printf("GetAbsPatherr:%v\n", err)
	}

	if records, err = zkteco.GetCSVRecords(p); err != nil {
		fmt.Printf("records:%v\n, err:%v\n", records, err)

	}
}
