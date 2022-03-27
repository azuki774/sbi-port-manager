package dbclient

import (
	"fmt"
	"os"

	csvimport "sbi-port-manager/internal/csv-import"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
)

var (
	Address   string
	Token     string
	Org       string
	Bucket    string
	FileNames []string
)

func getOptionFromEnv() {
	Address = os.Getenv("DB_ADDRESS")
	Token = os.Getenv("DB_TOKEN")
	Org = os.Getenv("DB_ORG")
	Bucket = os.Getenv("DB_BUCKET")
}

func PostMain(dir string) (err error) {
	getOptionFromEnv()
	client := influxdb.NewClientWithOptions(Address, Token, influxdb.DefaultOptions().SetBatchSize(20))
	writeAPI := client.WriteAPI(Org, Bucket)
	defer client.Close()

	filePaths, err := getFilePaths(dir)
	if err != nil {
		return err
	}

	for _, filePath := range filePaths { // １日分
		t, err := pathToDate(filePath)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		osf, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error: os.Open - %s\n", filePath)
			return err
		}

		records, err := csvimport.PortCSVToString(osf)
		if err != nil {
			fmt.Printf("Error: PortCSVToString - %s\n", filePath)
			return err
		}

		fundInfos, err := csvimport.FundsLoad(records)
		if err != nil {
			fmt.Printf("Error: FundsLoad - %s\n", filePath)
			return err
		}

		fmt.Println(filePath)

		for _, fundInfo := range fundInfos { // 1レコード分
			p := influxdb.NewPointWithMeasurement("port").
				AddTag("name", fundInfo.Name).
				AddField("count", fundInfo.Count).
				AddField("purchaseUnitPrice", fundInfo.PurchaseUnitPrice).
				AddField("nowPrice", fundInfo.NowPrice).
				AddField("beforeChange", fundInfo.BeforeChange).
				AddField("beforeRatio", fundInfo.BeforeRatio).
				AddField("profit", fundInfo.Profit).
				AddField("profitRatio", fundInfo.ProfitRatio).
				AddField("appraisedValue", fundInfo.AppraisedValue).
				SetTime(t)
			writeAPI.WritePoint(p)
		}
	}

	return nil
}
