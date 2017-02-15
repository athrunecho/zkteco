package main

import (
	"fmt"

	"github.com/AthrunEcho/zkteco"
	"github.com/northbright/pathhelper"
)

var (
	err error
	p   string
)

func main() {
	if p, err = pathhelper.GetAbsPath(kaoqingFilePath); err != nil {
		i := zkteco.GetCSVRecords(p)
		fmt.Printf("%v\n", i)
	}
}
