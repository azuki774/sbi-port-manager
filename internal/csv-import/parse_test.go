package csvimport

import (
	"reflect"
	"testing"
)

func Test_fundLoad(t *testing.T) {
	rowData1 := []string{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "}
	rowData2 := []string{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 ", "?????"} // too many elements
	rowData3 := []string{"積立  売却", "AAA", "--/--/--", "26231", "1.23456789", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "}     // not int

	fundInfo1 := fundInfo{}
	fundInfo2 := fundInfo{}
	fundInfo3 := fundInfo{}

	type args struct {
		fundInfo *fundInfo
		rowData  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{fundInfo: &fundInfo1, rowData: rowData1},
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{fundInfo: &fundInfo2, rowData: rowData2},
			wantErr: true,
		},
		{
			name:    "test3",
			args:    args{fundInfo: &fundInfo3, rowData: rowData3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fundLoad(tt.args.fundInfo, tt.args.rowData); (err != nil) != tt.wantErr {
				t.Errorf("fundLoad() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fundsLoad(t *testing.T) {
	rowData10 := []string{"取引", "ファンド名", "買付日", "数量", "取得単価", "現在値", "前日比", "前日比（％）", "損益", "損益（％）", "評価額", "編集"}
	rowData11 := []string{"積立  売却", "AAA", "--/--/--", "26231", "13000", "11403", "-258", "-2.21", "-4189.09", "-12.28", "29911.2", "詳細 "}
	rowData12 := []string{"積立  売却", "BBB", "--/--/--", "10946", "31610", "29726", "+235", "+0.80", "-2062.22", "-5.96", "32538.07", "詳細 "}
	csvData1 := [][]string{rowData10, rowData11, rowData12}

	collect1 := []fundInfo{
		{
			Name:              "AAA",    // ファンド名
			Count:             26231,    // 数量
			PurchaseUnitPrice: 13000,    // 取得単価
			NowPrice:          11403,    // 現在値
			BeforeChange:      -258,     // 前日比
			BeforeRatio:       -2.21,    // 前日比（％）
			Profit:            -4189.09, // 損益
			ProfitRatio:       -12.28,   // 損益（％）
			AppraisedValue:    29911.2,  // 評価額
		},
		{
			Name:              "BBB",    // ファンド名
			Count:             10946,    // 数量
			PurchaseUnitPrice: 31610,    // 取得単価
			NowPrice:          29726,    // 現在値
			BeforeChange:      235,      // 前日比
			BeforeRatio:       0.80,     // 前日比（％）
			Profit:            -2062.22, // 損益
			ProfitRatio:       -5.96,    // 損益（％）
			AppraisedValue:    32538.07, // 評価額
		},
	}

	type args struct {
		csvData [][]string
	}
	tests := []struct {
		name          string
		args          args
		wantFundsInfo []fundInfo
		wantErr       bool
	}{
		{
			name:          "test1",
			args:          args{csvData: csvData1},
			wantFundsInfo: collect1,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFundsInfo, err := fundsLoad(tt.args.csvData)
			if (err != nil) != tt.wantErr {
				t.Errorf("fundsLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFundsInfo, tt.wantFundsInfo) {
				t.Errorf("fundsLoad() = %v, want %v", gotFundsInfo, tt.wantFundsInfo)
			}
		})
	}
}
