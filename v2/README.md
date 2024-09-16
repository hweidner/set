[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDocs](https://godocs.io/github.com/hweidner/set/v2?status.svg)](https://godocs.io/github.com/hweidner/set/v2)
[![Go Reference](https://pkg.go.dev/badge/github.com/hweidner/set/v2.svg)](https://pkg.go.dev/github.com/hweidner/set/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/hweidner/set/v2)](https://goreportcard.com/report/github.com/hweidner/set/v2)

# set

A set implementation for Go using generics.

See http://en.wikipedia.org/wiki/Set_%28mathematics%29 for a full discussion
of sets.

This module contains two set implementations:

  * A generic implementation using interface{} values, for all versions of Go.
    This is fixed to v1.0.0. No more updates will happen.
  * An type-safe implementation using generics and iterator functions in the v2 directory.
    This requires Go 1.23 or later.

This documentation is for the version 2 of the package. See the
[README.md](../README.md) at the repositories top level for the documentation
of the version 1 package.

## Implementation

A set is implemented as map without values. The elements of a set are stored as maps keys.

## Examples

	import "github.com/hweidner/set/v2"
	
	a := set.New(1, 3, 5, 7, 9)
	b := set.New(2, 4, 6, 8, 10)

	fmt.Println(a, b)

	union := a.Union(b)
	intersect := a.Intersect(b)
	difference := a.Diff(b)

	fmt.Println(union, intersect, difference)

	a.Add(2)
	b.Add(5)
	fmt.Println(a.Intersect(b).Contains(2))
	
	for x := range a.All() {
		fmt.Println(x)
	}

## Rationale

All operations are invoked in an object oriented style. Only the Add, Remove
and Clear methods modify the receiver.

## License

This package is released under the MIT license.
The full license text can be found in the [LICENSE](LICENSE) file.
