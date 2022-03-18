package csvimport

import (
	"os"
	"reflect"
	"testing"
)

func Test_portCSVToString(t *testing.T) {
	type args struct {
		osf *os.File
	}
	tests := []struct {
		name        string
		args        args
		wantRecords [][]string
		wantErr     bool
	}{
		// TODO: Add test cases.
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
