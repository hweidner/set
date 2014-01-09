set
===

A generic set implementation for Go

Package set offeres a generic set implementation in Go.
See http://en.wikipedia.org/wiki/Set_%28mathematics%29 for a full discussion
of sets.

This package implements a set as a map without values. Keys can have any
arbitrary type, as long as there is equality defined on it.

The package documentation can be read on godoc.org:  
[![GoDoc](http://godoc.org/github.com/hweidner/set?status.png)](http://godoc.org/github.com/hweidner/set)

Examples
--------

	a := set.NewInit(1, 3, 5, 7, 9)
	b := set.NewInit(2, 4, 6, 8, 10)

	fmt.Println(a, b)

	union := a.Union(b)
	intersect := a.Intersect(b)
	difference := a.Diff(b)

	fmt.Println(union, intersect, difference)

	a.Add(2)
	b.Add(5)
	fmt.Println(a.Intersect(b).Contains(2))

Rationale
---------

All operations are invoked in an object oriented style. Only the Add, Remove
and Clear methods modify the receiver.

License
-------

This package is released under the GNU Lesser General Public License, Version
3. The full license text can be found in the LICENSE file of the source code
distribution.
