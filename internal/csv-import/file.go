package csvimport

import (
	"encoding/csv"
	"io"
	"os"
)

func portCSVToString(osf *os.File) (records [][]string, err error) {
	r := csv.NewReader(osf)
	index := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if index != 0 {
			// ラベルレコードは読み込まない
			records = append(records, record)
		}

		index++
	}
	return records, nil
}
