package zkteco

import (
	"os"
)

func GetCSVRecords(f string) (records [][]string, err error) {
	f, err = os.Open(f)
	if err != nil {
		debugPrintf("os.Open() err: %v\n", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err = r.ReadAll()
	if err != nil {
		debugPrintf("GetCSVRecords error: %v\n", err)
		return err
	}
}
