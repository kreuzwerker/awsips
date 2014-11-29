package awsips

import "sort"

type sortedSet map[string]struct{}

func newSortedSet(elements ...string) sortedSet {

	var set sortedSet = make(map[string]struct{})

	for _, e := range elements {
		set.Add(e)
	}

	return set

}

func (s sortedSet) Add(element string) {
	s[element] = struct{}{}
}

func (s sortedSet) All() []string {

	var set []string

	for k, _ := range s {
		set = append(set, k)
	}

	return set

}

func (s sortedSet) IsEmpty() bool {
	return len(s) == 0
}

func (s sortedSet) IsIncluded(element string) bool {

	for k, _ := range s {

		if k == element {
			return true
		}

	}

	return false

}

func (s sortedSet) Sorted() []string {

	var set sort.StringSlice = s.All()
	set.Sort()

	return set

}
