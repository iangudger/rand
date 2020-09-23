# `math/rand` compatible pseudo-random number generators powered by `crypto/rand`

A solution to golang.org/issue/25531.

`crypto/rand` is the secure source of random numbers in Go (golang), but it can be difficult to use. `math/rand` is much easier to use, but lacks the randomness properties of `crypto/rand`. This package tries to fix that problem.
