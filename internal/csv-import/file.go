package csvimport

import (
	"encoding/csv"
	"io"
	"os"
)

func PortCSVToString(osf *os.File) (records [][]string, err error) {
	r := csv.NewReader(osf)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}
	return records, nil
}
