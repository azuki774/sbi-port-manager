package dbclient

import (
	"reflect"
	"testing"
	"time"
)

func Test_pathToDate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantT   time.Time
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{path: "./test/afaf/1999-03-23.csv"},
			wantT:   time.Date(1999, 3, 23, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{path: "/afaf/2000-01-01.csv"},
			wantT:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "test3 (not csv)",
			args:    args{path: "./test/afaf/abc.csv"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, err := pathToDate(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("pathToDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("pathToDate() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_getFilePaths(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "internal/csv-import/test/",
			args:    args{dir: "../csv-import/test/"},
			want:    []string{"../csv-import/test/1999-01-23.csv", "../csv-import/test/1999-01-24.csv"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFilePaths(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFilePaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFilePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
