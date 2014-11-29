package awsips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func _result(t *testing.T) *Result {

	result, err := Client()
	assert.NoError(t, err)

	return result

}

func TestRegions(t *testing.T) {

	t.Parallel()

	regions := newSortedSet(_result(t).Regions()...)

	for _, region := range []string{
		"GLOBAL",
		"ap-northeast-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"cn-north-1",
		"eu-central-1",
		"eu-west-1",
		"sa-east-1",
		"us-east-1",
		"us-gov-west-1",
		"us-west-1",
		"us-west-2",
	} {
		assert.True(t, regions.IsIncluded(region))
	}

}

func TestServices(t *testing.T) {

	t.Parallel()

	services := newSortedSet(_result(t).Services()...)

	for _, service := range []string{
		"AMAZON",
		"CLOUDFRONT",
		"EC2",
		"ROUTE53",
		"ROUTE53_HEALTHCHECKS",
	} {
		assert.True(t, services.IsIncluded(service))
	}

}

func TestFilter(t *testing.T) {

	t.Parallel()

	assert := assert.New(t)
	noFilter := []string{}
	ipv4Cidr := `^(?:\d{1,3}\.){3}\d{1,3}\/\d+$`

	result, err := Client()

	assert.NoError(err)

	tree := result.Filter(noFilter, noFilter)

	assert.NotNil(tree["CLOUDFRONT"])
	assert.Nil(tree["SILVER_BULLET_SERVICE"])

	assert.Regexp(ipv4Cidr, tree["CLOUDFRONT"]["GLOBAL"][0])

	tree = result.Filter(noFilter, []string{"EC2", "AMAZON"})

	assert.Nil(tree["CLOUDFRONT"])
	assert.NotNil(tree["AMAZON"])
	assert.NotNil(tree["EC2"])

	assert.NotNil(tree["EC2"]["eu-west-1"])
	assert.NotNil(tree["EC2"]["eu-central-1"])
	assert.NotNil(tree["EC2"]["us-east-1"])

	assert.Regexp(ipv4Cidr, tree["EC2"]["eu-west-1"][0])

	tree = result.Filter([]string{"eu-west-1", "eu-central-1"}, []string{"EC2", "AMAZON"})

	assert.Nil(tree["CLOUDFRONT"])
	assert.NotNil(tree["AMAZON"])
	assert.NotNil(tree["EC2"])

	assert.NotNil(tree["EC2"]["eu-west-1"])
	assert.NotNil(tree["EC2"]["eu-central-1"])
	assert.Nil(tree["EC2"]["us-east-1"])

	assert.Regexp(ipv4Cidr, tree["EC2"]["eu-west-1"][0])

}
