// Copyright 2014 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the GNU Lesser General Public License
// Version 3 that can be found in the LICENSE file.

package set

import (
	"fmt"
	"sort"
)

// ----- AnySlice -----

// AnySlice is a slice type for arbitrary values, which implements the sort.Interface.
// We use it for returning a list of (sorted or unsorted) set elements.
type AnySlice []interface{}

// The Len function for sort.Interface.
func (s AnySlice) Len() int {
	return len(s)
}

// The Less function for sort.Interface.
func (s AnySlice) Less(i, j int) bool {
	return fmt.Sprint(s[i]) < fmt.Sprint(s[j])
}

// The Swap function for sort.Interface.
func (s AnySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// ----- Set definition -----

// A set is implemented as a map without values.
type Set map[interface{}]struct{}

// how to initialize the map entry
var value = struct{}{}

// ----- constructors -----

// Create a new set.
func New() Set {
	return Set{}
}

// Create a new initialized set.
func NewInit(e ...interface{}) Set {
	s := Set{}
	for _, i := range e {
		s[i] = value
	}
	return s
}

// ----- methods -----

// IsEmpty tests if the set is empty.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Len returns the length of the set.
func (s Set) Len() int {
	return len(s)
}

// Contains checks if a set contains one or more elements. The return value
// is true only if all given elements are in the set.
func (s Set) Contains(e ...interface{}) bool {
	for _, i := range e {
		if _, ok := s[i]; !ok {
			return false
		}
	}
	return true
}

// IsEqual tests if two sets are equal.
func (s Set) IsEqual(t Set) bool {
	if len(s) != len(t) {
		return false
	}
	for k, _ := range s {
		if _, ok := t[k]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf returns true if the set s is a subset of the set t, e.g. if
// all elements of s are also in t.
func (s Set) IsSubsetOf(t Set) bool {
	for k, _ := range s {
		if _, ok := t[k]; !ok {
			return false
		}
	}
	return true
}

// IsSupersetOf returns true if the set s is a superset of the set t, e.g.
// if all elements of t are also in s.
func (s Set) IsSupersetOf(t Set) bool {
	return t.IsSubsetOf(s)
}

// Clear removes all elements from the given set.
func (s Set) Clear() {
	for k, _ := range s {
		delete(s, k)
	}
}

// Copy returns a copy of a set. The set s is not modified.
func (s Set) Copy() Set {
	t := Set{}
	for k, _ := range s {
		t[k] = value
	}
	return t
}

// Add adds one or more elements to the given set.
func (s Set) Add(e ...interface{}) {
	for _, i := range e {
		s[i] = value
	}
}

// Remove removes one or more elements from the given set.
func (s Set) Remove(e ...interface{}) {
	for _, i := range e {
		delete(s, i)
	}
}

// Union returns a new set, which represents the union of two or more sets.
// The sets themselfes are not modified.
func (s Set) Union(t ...Set) Set {
	r := Set{}
	for k, _ := range s {
		r[k] = value
	}
	for _, i := range t {
		for k, _ := range i {
			r[k] = value
		}
	}
	return r
}

// Intersect returns a new set which represents the intersection of two or more sets.
// The sets themselfes are not modified.
func (s Set) Intersect(t ...Set) Set {
	r := Set{}
	for k, _ := range s {
		inter := true
		for _, i := range t {
			if _, ok := i[k]; !ok {
				inter = false
				break
			}
		}
		if inter == true {
			r[k] = value
		}
	}
	return r
}

// Diff returns a new set which represents the difference of two sets.
// The sets themselfes are not modified.
func (s Set) Diff(t Set) Set {
	r := Set{}
	for k, _ := range s {
		if _, ok := t[k]; !ok {
			r[k] = value
		}
	}
	return r
}

// SymDiff returns a new set which represents the symmetric difference of two
// sets. The sets themselfes are not modified.
func (s Set) SymDiff(t Set) Set {
	r := s.Copy()
	for k, _ := range t {
		if _, ok := s[k]; !ok {
			r[k] = value
		} else {
			delete(r, k)
		}
	}
	return r
}

// List returns a list of the set elements in a slice.
func (s Set) List() AnySlice {
	r := make(AnySlice, 0, len(s))
	for k, _ := range s {
		r = append(r, k)
	}
	return r
}

// List returns a list of the set elements in a slice.
func (s Set) SortedList() AnySlice {
	r := s.List()
	sort.Sort(r)
	return r
}

// Function for implementing the fmt.Stringer interface to prettyprint the set.
func (s Set) String() string {
	str := "{ "
	for k, _ := range s {
		str += fmt.Sprint(k) + " "
	}
	str += "}"
	return str
}
