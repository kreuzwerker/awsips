package awsips

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	filter    = ""
	result, _ = Client()
	output    = &Output{result}
)

func TestFormat(t *testing.T) {

	assert := assert.New(t)

	out, err := format([]string{"foo", "bar "})

	assert.NoError(err)
	assert.Equal(out, `[
    "foo",
    "bar "
]`)

}

func TestId(t *testing.T) {

	var tt = []struct {
		service string
		region  string
		out     string
	}{
		{"ROUTE53_HEALTHCHECKS", "ap-northeast-1", "route53HealthchecksApNortheast1"},
	}

	for _, e := range tt {
		assert.Equal(t, e.out, id(e.service, e.region))
	}

}

func TestSplit(t *testing.T) {

	assert := assert.New(t)

	a := "foo,bar,wee"
	b := ""

	assert.Equal([]string{"foo", "bar", "wee"}, split(&a))
	assert.Empty(split(&b))

}

func TestOutputFiles(t *testing.T) {

	assert := assert.New(t)

	dir, err := ioutil.TempDir("", "")
	defer os.Remove(dir)

	assert.NoError(err)

	files, err := ioutil.ReadDir(dir)

	assert.NoError(err)
	assert.Empty(files)

	err = os.Chdir(dir)

	assert.NoError(err)

	err = output.Files(&filter, &filter)

	files, err = ioutil.ReadDir(dir)

	assert.NoError(err)
	assert.NotEmpty(files)

	out, err := ioutil.ReadFile(filepath.Join(dir, "cloudfrontGlobal.json"))

	assert.NoError(err)
	assert.Contains(string(out), `[
    "204.246.164.0/22",`)

}

func TestOutputRegions(t *testing.T) {

	assert := assert.New(t)

	out, err := output.Regions()

	assert.NoError(err)
	assert.Contains(out, `"GLOBAL",`)

}

func TestOutputServices(t *testing.T) {

	assert := assert.New(t)

	out, err := output.Services()

	assert.NoError(err)
	assert.Contains(out, `"CLOUDFRONT",`)

}

func TestOutputStack(t *testing.T) {

	assert := assert.New(t)
	ports := ""

	out, err := output.Stack(&filter, &filter, &ports)

	assert.NoError(err)
	assert.Contains(out, `                        "FromPort": "1",
                        "IpProtocol": "tcp",
                        "ToPort": "65535"
                    },`)

	ports = "80,443"

	out, err = output.Stack(&filter, &filter, &ports)

	assert.NoError(err)
	assert.Contains(out, `                        "FromPort": "80",
                        "IpProtocol": "tcp",
                        "ToPort": "80"
                    },`)
	assert.Contains(out, `                        "FromPort": "443",
                        "IpProtocol": "tcp",
                        "ToPort": "443"
                    },`)

}

func TestOutputTree(t *testing.T) {

	assert := assert.New(t)

	out, err := output.Tree(&filter, &filter)

	assert.NoError(err)
	assert.Contains(out, `{
    "AMAZON": {
        "GLOBAL": [`)

}
