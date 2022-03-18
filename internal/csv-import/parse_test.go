package csvimport

import "testing"

func Test_fundLoad(t *testing.T) {
	type args struct {
		fundInfo *fundInfo
		rowData  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := fundLoad(tt.args.fundInfo, tt.args.rowData); (err != nil) != tt.wantErr {
				t.Errorf("fundLoad() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
