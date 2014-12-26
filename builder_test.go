package awsips

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {

	assert := assert.New(t)

	root := N()

	root.L("Integer", 1)
	root.L("String", "Hello world")
	root.N("Node").L("foo", "bar")

	out := `{
    "Integer": 1,
    "Node": {
        "foo": "bar"
    },
    "String": "Hello world"
}`

	m, err := json.MarshalIndent(root, "", "    ")

	assert.NoError(err)
	assert.Equal(out, string(m))

}
