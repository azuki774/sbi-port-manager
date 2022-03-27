package dbclient

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

// csvのpathから日付を取得。時刻は設定しないのでUTC 00:00になる
func pathToDate(path string) (t time.Time, err error) {
	fileName := filepath.Base(path)                          // To FileName yyyy-mm-dd.csv
	basefileName := strings.ReplaceAll(fileName, ".csv", "") // yyyy-mm-dd

	layout := "2006-01-02"
	t, err = time.Parse(layout, basefileName)
	if err != nil {
		return t, fmt.Errorf("pathToDate : Invalid FileName")
	}

	return t, nil
}

func getFilePaths(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}
