package zkteco

import (
	"encoding/csv"
	"io"
)

func GetCSVRecords(r io.Reader) (records [][]string, err error) {
	reader := csv.NewReader(r)
	records, err = reader.ReadAll()
	if err != nil {
		debugPrintf("GetCSVRecords error: %v\n", err)
		return [][]string{}, err
	}
	return records, nil
}
