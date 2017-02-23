package zkteco_test

import (
	"log"
	"os"

	"github.com/athrunecho/zkteco"
	"github.com/northbright/pathhelper"
)

const (
	kaoqinFilePath = "./example/csv/attendance/attendance.csv"
)

// 1. Run "go test -c && ./zkteco.test" to generate test binary under root of project dir and start test.
//
// 2. Run "go test" will generate test binary under "/tmp" with a random test dir
func ExampleGetCSVRecords() {
	var (
		err error
		p   string
		f   *os.File
		re  [][]string
	)

	log.Printf("\n")
	log.Printf("--------- GetCSVRecords() Test Begin --------\n")

	if p, err = pathhelper.GetAbsPath(kaoqinFilePath); err != nil {
		goto end
	}

	if f, err = os.Open(p); err != nil {
		goto end
	}
	defer f.Close()

	if re, err = zkteco.GetCSVRecords(f); err != nil {
		goto end
	}

	log.Printf("%v\n", re)

end:
	if err != nil {
		log.Printf("err: %v\n", err)
	}
	log.Printf("--------- BatchCreate() Test End --------\n")
	// Output:
}
