package common

import "strings"

type StringSet struct {
	paths map[string]struct{}
}

func (set *StringSet) Set(s string) {
	set.checkNil()
	set.paths[s] = struct{}{}
}

func (set *StringSet) All() (keys []string) {
	set.checkNil()
	for key := range set.paths {
		keys = append(keys, key)
	}
	return
}

func (set *StringSet) Size() int {
	set.checkNil()
	return len(set.paths)
}

func (set *StringSet) checkNil() {
	if set.paths == nil {
		set.paths = make(map[string]struct{})
	}
}

func (set *StringSet) String() string {
	return strings.Join(set.All(), ",")
}

func (set *StringSet) SetAll(ss []string) {
	set.checkNil()
	for _, s := range ss {
		set.Set(s)
	}
}

func (set *StringSet) Contains(s string) bool {
	set.checkNil()
	_, ok := set.paths[s]
	return ok
}

func (set *StringSet) Minus(other *StringSet) StringSet {

	set.checkNil()
	ret := StringSet{}

	for s := range set.paths {
		if !other.Contains(s) {
			ret.Set(s)
		}
	}

	return ret
}
