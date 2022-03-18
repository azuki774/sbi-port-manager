package csvimport

import (
	"fmt"
	"strconv"
)

const (
	csvElementSize = 12
)

func fundsLoad(csvData [][]string) (fundsInfo []fundInfo, err error) {
	index := 0
	for _, v := range csvData {
		var nowfundInfo fundInfo
		err := fundLoad(&nowfundInfo, v)
		if err != nil {
			return nil, err
		}
		if index != 0 {
			// ラベル部分は取り込まない
			fundsInfo = append(fundsInfo, nowfundInfo)
		}
		index++
	}
	return fundsInfo, nil
}

func fundLoad(fundInfo *fundInfo, rowData []string) (err error) {
	if len(rowData) != csvElementSize {
		return fmt.Errorf("fundLoad parse error")
	}

	fundInfo.Name = rowData[1] // ファンド名

	fundInfo.Count, err = strconv.Atoi(rowData[3]) // 数量
	if err != nil {
		return fmt.Errorf("fundInfo.Count Atoi error")
	}

	fundInfo.PurchaseUnitPrice, err = strconv.Atoi(rowData[4]) // 取得単価
	if err != nil {
		return fmt.Errorf("fundInfo.PurchaseUnitPrice Atoi error")
	}

	fundInfo.NowPrice, err = strconv.Atoi(rowData[5]) // 現在値
	if err != nil {
		return fmt.Errorf("fundInfo.NowPrice Atoi error")
	}

	fundInfo.BeforeChange, err = strconv.Atoi(rowData[6]) // 前日比
	if err != nil {
		return fmt.Errorf("fundInfo.BeforeChange Atoi error")
	}

	fundInfo.BeforeRatio, err = strconv.ParseFloat(rowData[7], 64) // 前日比（％）
	if err != nil {
		return fmt.Errorf("fundInfo.BeforeRatio ParseFloat error")
	}

	fundInfo.Profit, err = strconv.ParseFloat(rowData[8], 64) // 損益
	if err != nil {
		return fmt.Errorf("fundInfo.Profit ParseFloat error")
	}

	fundInfo.ProfitRatio, err = strconv.ParseFloat(rowData[9], 64) // 損益（％）
	if err != nil {
		return fmt.Errorf("fundInfo.ProfitRatio ParseFloat error")
	}

	fundInfo.AppraisedValue, err = strconv.ParseFloat(rowData[10], 64) // 評価額
	if err != nil {
		return fmt.Errorf("fundInfo.AppraisedValue ParseFloat error")
	}

	return nil
}
