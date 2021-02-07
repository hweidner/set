// Copyright 2014 by Harald Weidner <hweidner@gmx.net>. All rights reserved.
// Use of this source code is governed by the MIT license. See the LICENSE file
// for a full text of the license.
// SPDX-License-Identifier: MIT

package set

import (
	"fmt"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	a := NewInit(0, 1, 3, 5, 7, 9)
	b := NewInit(0, 2, 4, 6, 8, 10)
	zero := New()
	zero.Add(0)
	full := NewInit(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	aStr := "[0 1 3 5 7 9]"

	un := a.Union(b)
	if !un.IsEqual(full) {
		t.Errorf("Union failed: expected %v, got %v.\n", full, un)
	}

	in := a.Intersect(b)
	if !in.IsEqual(zero) {
		t.Errorf("Intersect failed: expected %v, got %v.\n", zero, in)
	}

	str := fmt.Sprint(a.SortedList())
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

	e := un.SortedList()
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

	ch, done := a.Iterator()
	h1 := <-ch
	h2 := <-ch
	close(done)
	time.Sleep(10 * time.Millisecond)
	h3, ok := <-ch
	if h1 == interface{}(nil) || h2 == interface{}(nil) || h3 != interface{}(nil) || ok != false {
		t.Errorf("Iterator failed: got %T/%T/%T/%t, expected int/int/<nil>/false", h1, h2, h3, ok)
	}
}
