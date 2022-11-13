package cache

import (
	"path/filepath"
	"testing"

	"github.com/Kindred87/Spoke/internal/common"
	"github.com/Kindred87/Spoke/internal/io"
	"github.com/stretchr/testify/assert"
)

var (
	testFile = filepath.Join("..", "..", "testdata", common.TestFile)
)

func TestBuildFrom(t *testing.T) {
	tree, err := io.Load(testFile)
	assert.Nil(t, err)

	type args struct {
		tree map[string]any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Basic", args: args{tree: tree}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BuildFrom(tt.args.tree); (err != nil) != tt.wantErr {
				t.Errorf("BuildFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
