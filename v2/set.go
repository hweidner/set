// Copyright 2014-2022 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

/*
A set library for Go using generics

Package set provides a generic type-safe set library in Go, using the new generics
language extension in Go 1.18 and higher.

*/
package set

import "fmt"

// ----- Set definition -----

// The Set is implemented as a map without values.
type Set[T comparable] struct {
	set map[T]struct{}
}

// ----- constructor -----

// New creates a new set and initializes it with the argument values.
// Note: NewInit() was renamed to New()
func New[T comparable](e ...T) Set[T] {
	s := Set[T]{set: make(map[T]struct{}, len(e))}
	for _, i := range e {
		s.set[i] = struct{}{}
	}
	return s
}

// ----- methods that modify the receiver -----

// Add adds one or more elements to the given set.
func (s Set[T]) Add(e ...T) {
	for _, i := range e {
		s.set[i] = struct{}{}
	}
}

// Remove removes one or more elements from the given set.
func (s Set[T]) Remove(e ...T) {
	for _, i := range e {
		delete(s.set, i)
	}
}

// Clear removes all elements from the given set.
func (s Set[T]) Clear() {
	for k := range s.set {
		delete(s.set, k)
	}
}

// ----- methods that do not modify the receiver -----

// IsEmpty tests if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s.set) == 0
}

// Len returns the length of the set.
func (s Set[T]) Len() int {
	return len(s.set)
}

// Contains checks if a set contains one or more elements. The return value
// is true only if all given elements are in the set.
func (s Set[T]) Contains(e ...T) bool {
	for _, i := range e {
		if _, ok := s.set[i]; !ok {
			return false
		}
	}
	return true
}

// IsEqual tests if two sets are equal.
func (s Set[T]) IsEqual(t Set[T]) bool {
	if len(s.set) != len(t.set) {
		return false
	}
	for k := range s.set {
		if _, ok := t.set[k]; !ok {
			return false
		}
	}
	return true
}

// IsSubsetOf returns true if the set s is a subset of the set t, e.g. if
// all elements of s are also in t.
func (s Set[T]) IsSubsetOf(t Set[T]) bool {
	for k := range s.set {
		if _, ok := t.set[k]; !ok {
			return false
		}
	}
	return true
}

// IsSupersetOf returns true if the set s is a superset of the set t, e.g.
// if all elements of t are also in s.
func (s Set[T]) IsSupersetOf(t Set[T]) bool {
	return t.IsSubsetOf(s)
}

// ----- methods that return a new set -----

// Copy returns a copy of a set. The set s is not modified.
func (s Set[T]) Copy() Set[T] {
	r := Set[T]{set: make(map[T]struct{}, len(s.set))}
	for k := range s.set {
		r.set[k] = struct{}{}
	}
	return r
}

// Union returns a new set, which represents the union of two or more sets.
// The sets themselves are not modified.
func (s Set[T]) Union(t ...Set[T]) Set[T] {
	// calculate overall length of sets
	l := len(s.set)
	for _, i := range t {
		l += len(i.set)
	}

	// create result set. As a heuristic, the estimated length is 50% of the sum of the lengths
	// of each input set.
	r := Set[T]{set: make(map[T]struct{}, l/2)}

	for k := range s.set {
		r.set[k] = struct{}{}
	}
	for _, i := range t {
		for k := range i.set {
			r.set[k] = struct{}{}
		}
	}
	return r
}

// Intersect returns a new set which represents the intersection of two or more sets.
// The sets themselves are not modified.
func (s Set[T]) Intersect(t ...Set[T]) Set[T] {
	r := Set[T]{set: make(map[T]struct{}, len(s.set))}
next_s_elem:
	for k := range s.set {
		for _, i := range t {
			if _, ok := i.set[k]; !ok {
				continue next_s_elem
			}
		}
		r.set[k] = struct{}{}
	}
	return r
}

// Diff returns a new set which represents the difference of two sets.
// The sets themselves are not modified.
func (s Set[T]) Diff(t Set[T]) Set[T] {
	r := Set[T]{set: make(map[T]struct{}, len(s.set))}
	for k := range s.set {
		if _, ok := t.set[k]; !ok {
			r.set[k] = struct{}{}
		}
	}
	return r
}

// SymDiff returns a new set which represents the symmetric difference of two
// sets. The sets themselves are not modified.
func (s Set[T]) SymDiff(t Set[T]) Set[T] {
	r := s.Copy()
	for k := range t.set {
		if _, ok := s.set[k]; !ok {
			r.set[k] = struct{}{}
		} else {
			delete(r.set, k)
		}
	}
	return r
}

// ----- methods that return other data types -----

// List returns an unsorted list of the set elements in a slice.
func (s Set[T]) List() []T {
	r := make([]T, 0, len(s.set))

	for k := range s.set {
		r = append(r, k)
	}
	return r
}

// Iterator returns a channel that can be used to iterate over the set. A second
// "done" channel can be used to preliminarily terminate the iteration by closing
// the done channel.
func (s Set[T]) Iterator() (<-chan T, chan<- struct{}) {
	ic := make(chan T)
	done := make(chan struct{})
	go func() {
		for k := range s.set {
			select {
			case ic <- k:
			case <-done:
				break
			}
		}
		close(ic)
	}()
	return ic, done
}

// String returns a textual representation of the set in a string.
// It is there for implementing the fmt.Stringer interface to prettyprint the set.
func (s Set[T]) String() string {
	str := "{ "
	for k := range s.set {
		str += fmt.Sprint(k) + " "
	}
	str += "}"
	return str
}
