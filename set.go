// Copyright 2014 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

package set

import (
	"fmt"
	"sort"
)

// VERSION is the version number of the set package
const VERSION = "0.1"

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

// The Set is implemented as a map without values.
type Set map[interface{}]struct{}

// how to initialize the map entry
var value = struct{}{}

// ----- constructors -----

// New creates a new (empty) set.
func New() Set {
	return Set{}
}

// NewInit creates a new set and initializes it with the argument values.
func NewInit(e ...interface{}) Set {
	s := Set{}
	for _, i := range e {
		s[i] = value
	}
	return s
}

// ----- methods that modify the receiver -----

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

// Clear removes all elements from the given set.
func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// ----- methods that do not modify the receiver -----

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
	for k := range s {
		if _, ok := t[k]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf returns true if the set s is a subset of the set t, e.g. if
// all elements of s are also in t.
func (s Set) IsSubsetOf(t Set) bool {
	for k := range s {
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

// Copy returns a copy of a set. The set s is not modified.
func (s Set) Copy() Set {
	t := Set{}
	for k := range s {
		t[k] = value
	}
	return t
}

// Union returns a new set, which represents the union of two or more sets.
// The sets themselves are not modified.
func (s Set) Union(t ...Set) Set {
	r := Set{}
	for k := range s {
		r[k] = value
	}
	for _, i := range t {
		for k := range i {
			r[k] = value
		}
	}
	return r
}

// Intersect returns a new set which represents the intersection of two or more sets.
// The sets themselves are not modified.
func (s Set) Intersect(t ...Set) Set {
	r := Set{}
next_s_elem:
	for k := range s {
		for _, i := range t {
			if _, ok := i[k]; !ok {
				continue next_s_elem
			}
		}
		r[k] = value
	}
	return r
}

// Diff returns a new set which represents the difference of two sets.
// The sets themselves are not modified.
func (s Set) Diff(t Set) Set {
	r := Set{}
	for k := range s {
		if _, ok := t[k]; !ok {
			r[k] = value
		}
	}
	return r
}

// SymDiff returns a new set which represents the symmetric difference of two
// sets. The sets themselves are not modified.
func (s Set) SymDiff(t Set) Set {
	r := s.Copy()
	for k := range t {
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
	for k := range s {
		r = append(r, k)
	}
	return r
}

// SortedList returns a sorted list of the set elements in a slice.
func (s Set) SortedList() AnySlice {
	r := s.List()
	sort.Sort(r)
	return r
}

// Function for implementing the fmt.Stringer interface to prettyprint the set.
func (s Set) String() string {
	str := "{ "
	for k := range s {
		str += fmt.Sprint(k) + " "
	}
	str += "}"
	return str
}

// Iterator returns a channel that can be used to iterate over the set. A second
// "done" channel can be used to preliminarily terminate the iteration by closing
// the done channel.
func (s Set) Iterator() (<-chan interface{}, chan<- struct{}) {
	ic := make(chan interface{})
	done := make(chan struct{})
	go func() {
		for k := range s {
			select {
			case ic <- k:
			case <-done:
				close(ic)
				return
			}
		}
		close(ic)
	}()
	return ic, done
}
