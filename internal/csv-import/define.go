package csvimport

type fundInfo struct {
	Name              string  // ファンド名
	Count             int     // 数量
	PurchaseUnitPrice int     // 取得単価
	NowPrice          int     // 現在値
	BeforeChange      int     // 前日比
	BeforeRatio       float64 // 前日比（％）
	Profit            float64 // 損益
	ProfitRatio       float64 // 損益（％）
	AppraisedValue    float64 // 評価額
}

// func ImportMain() {
// 	f, err := os.Open("/home/azuki/work/sbi-port-manager/internal/csv-import/test/1999-01-23.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
