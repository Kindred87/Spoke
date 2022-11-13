package spoke

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_yml(t *testing.T) {
	fi, err := os.ReadFile("testdata/foo.yml")
	assert.Nil(t, err)

	ast := make(map[string]interface{})

	err = yaml.Unmarshal(fi, &ast)
	assert.Nil(t, err)

}
