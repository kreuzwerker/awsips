package awsips

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var Client = func() (*Result, error) {

	res, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var result Result

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil

}

type Result struct {
	CreateDate string
	Prefixes   []Prefix
	SyncToken  string
}

type Prefix struct {
	Ip_Prefix string
	Region    string
	Service   string
}

type ServiceTree map[string]RegionTree
type RegionTree map[string][]string

func (r *Result) Regions() []string {

	var set = newSortedSet()

	for _, prefix := range r.Prefixes {
		set.Add(prefix.Region)
	}

	return set.Sorted()

}

func (r *Result) Services() []string {

	var set = newSortedSet()

	for _, prefix := range r.Prefixes {
		set.Add(prefix.Service)
	}

	return set.Sorted()

}

func (r *Result) Filter(regions, services []string) ServiceTree {

	var (
		filterR = newSortedSet(regions...)
		filterS = newSortedSet(services...)
		group   = make(map[string]RegionTree)
	)

	for _, prefix := range r.Prefixes {

		var (
			r = filterR.IsEmpty() || filterR.IsIncluded(prefix.Region)
			s = filterS.IsEmpty() || filterS.IsIncluded(prefix.Service)
		)

		if r && s {

		retry:

			rg := group[prefix.Service]

			if rg == nil {
				group[prefix.Service] = make(map[string][]string)
				goto retry
			}

			set := newSortedSet(rg[prefix.Region]...)
			set.Add(prefix.Ip_Prefix)

			rg[prefix.Region] = set.Sorted()

		}

	}

	return group

}
