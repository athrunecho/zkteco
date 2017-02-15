package zkteco

import (
	"encoding/csv"
	"os"
)

func GetCSVRecords(file string) (records [][]string, err error) {
	var f *os.File

	f, err = os.Open(file)
	if err != nil {
		debugPrintf("os.Open() err: %v\n", err)
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err = r.ReadAll()
	if err != nil {
		debugPrintf("GetCSVRecords error: %v\n", err)
		return [][]string{}, err
	}
	return records, nil
}
