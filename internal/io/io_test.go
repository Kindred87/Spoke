package io

import (
	"path/filepath"
	"testing"

	"github.com/Kindred87/Spoke/internal/common"
)

var (
	testFile = filepath.Join("..", "..", "testdata", common.TestFile)
)

func Test_load(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Basic", args: args{file: testFile}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
