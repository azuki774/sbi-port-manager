package csvimport

import (
	"os"
	"reflect"
	"testing"
)

func Test_portCSVToString(t *testing.T) {
	dir, _ := os.Getwd()
	testfile1, _ := os.Open(dir + "/test/1999-01-23.csv")
	rowData10 := []string{"取引", "ファンド名", "買付日", "数量", "取得単価", "現在値", "前日比", "前日比（％）", "損益", "損益（％）", "評価額", "編集"}
	rowData11 := []string{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "}
	rowData12 := []string{"積立  売却", "BBB", "--/--/--", "10946", "31610", "29726", "+235", "+0.80", "-2062.22", "-5.96", "32538.07", "詳細 "}
	collect1 := [][]string{rowData10, rowData11, rowData12}

	type args struct {
		osf *os.File
	}
	tests := []struct {
		name        string
		args        args
		wantRecords [][]string
		wantErr     bool
	}{
		{
			name:        "test/1999-01-23.csv",
			args:        args{osf: testfile1},
			wantRecords: collect1,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := portCSVToString(tt.args.osf)
			if (err != nil) != tt.wantErr {
				t.Errorf("portCSVToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("portCSVToString() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
