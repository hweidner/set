// Copyright 2014-2022 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

package set

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func TestSet(t *testing.T) {
	a := New[int](0, 1, 3, 5, 7, 9)
	b := New[int](0, 2, 4, 6, 8, 10)
	zero := New[int]()
	zero.Add(0)
	full := New[int](0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	aStr := "[0 1 3 5 7 9]"

	un := a.Union(b)
	if !un.IsEqual(full) {
		t.Errorf("Union failed: expected %v, got %v.\n", full, un)
	}

	in := a.Intersect(b)
	if !in.IsEqual(zero) {
		t.Errorf("Intersect failed: expected %v, got %v.\n", zero, in)
	}

	l := a.List()
	slices.Sort(l)
	str := fmt.Sprint(l)
	if str != aStr {
		t.Errorf("SortedList failed: expected %v, got %v.\n", str, aStr)
	}

	if !a.IsSupersetOf(in) {
		t.Errorf("IsSupersectOf failed: %v is expected to be a superset of %v.\n", a, in)
	}
	if !b.IsSubsetOf(un) {
		t.Errorf("IsSubsetOf failed: %v is expected to be a subset of %v.\n", b, un)
	}
	if un.Len() != 11 || in.Len() != 1 {
		t.Errorf("Len Error: got union length %d (expected: 11) and intersect length %d (expected: 1).\n",
			un.Len(), in.Len())
	}

	if a.Contains(1, 5, 8) {
		t.Errorf("Contains failed: a should not contain value 8.\n")
	}
	if !b.Contains(8, 6, 4) {
		t.Errorf("Contains failed: b should contain values 8, 6, and 4.\n")
	}

	c := b.Copy()
	if !c.IsEqual(b) {
		t.Errorf("Copy(IsEqual failed: %v and %v expected to be equal.\n", c, b)
	}

	d := a.SymDiff(b)
	d.Add(0)
	if !un.IsEqual(d) {
		t.Errorf("Diff/Add/IsEqual failed: %v and %v were expected equal.\n", un, d)
	}

	e := un.List()
	slices.Sort(e)
	if len(e) != 11 || e[1] != 1 {
		t.Errorf("SortedList failed: length of %v should be 11 and the 2nd element 1.\n", e)
	}

	f := a.Diff(b)
	f.Add(0)
	if !a.IsEqual(f) {
		t.Errorf("Diff/Add/IsEqual failed: %v and %v should be equal.\n", a, f)

	}

	g := fmt.Sprint(zero)
	if g != "{ 0 }" {
		t.Errorf("String failed: got %v, expected \"{ 0 }\".\n", g)

	}

	zero.Remove(0)
	if zero.Len() != 0 {
		t.Errorf("Remove/Len failed: length of zero should be 0.\n")
	}

	b.Clear()
	if !b.IsEmpty() {
		t.Errorf("Clear/IsEmpty test failed: b.Clear() is not empty but %v.\n", b)
	}

	ch, done := New(1, 2, 3, 4, 5).Iterator()
	h1 := <-ch // must be 1..5
	h2 := <-ch // dto.
	close(done)
	time.Sleep(10 * time.Millisecond)
	h3, ok := <-ch // must be 0/false because iteration is done
	if h1 == 0 || h2 == 0 || h3 != 0 || ok != false {
		t.Errorf("Iterator failed: got %d/%d/%d/%t, expected 1-5/1-5/0/false", h1, h2, h3, ok)
	}
}
