[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDocs](https://godocs.io/github.com/hweidner/set?status.svg)](https://godocs.io/github.com/hweidner/set)
[![Go Reference](https://pkg.go.dev/badge/github.com/hweidner/set.svg)](https://pkg.go.dev/github.com/hweidner/set)
[![Go Report Card](https://goreportcard.com/badge/github.com/hweidner/set)](https://goreportcard.com/report/github.com/hweidner/set)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/hweidner/set.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/hweidner/set/alerts/)

set
===

A set implementation for Go using generics.

See http://en.wikipedia.org/wiki/Set_%28mathematics%29 for a full discussion
of sets.

This package implements a set as a map without values. Keys can have any
arbitrary type, as long as there is equality defined on it.

This module contains two set implementations:

  * A generic implementation using interface{} values, for all versions of Go.
    This is fixed to v1.0.0. No more updates will happen.
  * An type-safe implementation using generics in the v2 directory.
    This requires Go 1.18 or later.

This documentation is for the version 2 of the package. See the
[README.md](../README.md) at the repositories top level for the documentation
of the version 1 package.

Examples
--------

	import "github.com/hweidner/set/v2"
	
	a := set.New[int](1, 3, 5, 7, 9)
	b := set.New[int](2, 4, 6, 8, 10)

	fmt.Println(a, b)

	union := a.Union(b)
	intersect := a.Intersect(b)
	difference := a.Diff(b)

	fmt.Println(union, intersect, difference)

	a.Add(2)
	b.Add(5)
	fmt.Println(a.Intersect(b).Contains(2))
	
	ch, _ := a.Iterator()
	for x := range ch {
		fmt.Println(x)
	}


Rationale
---------

All operations are invoked in an object oriented style. Only the Add, Remove
and Clear methods modify the receiver.

License
-------

This package is released under the MIT license.
The full license text can be found in the [LICENSE](LICENSE) file.

