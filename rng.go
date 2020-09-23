// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"encoding/binary"
	"io"
)

type source struct {
	reader io.Reader
}

// Seed is stubbed out to implement math/rand.Source.
func (s *source) Seed(seed int64) {}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (s *source) Int63() int64 {
	const mask = (1 << 63) - 1
	return int64(s.Uint64() & mask)
}

// Uint64 returns a non-negative pseudo-random 64-bit integer as an uint64.
func (s *source) Uint64() uint64 {
	b := make([]byte, 8)
	if _, err := io.ReadFull(s.reader, b); err != nil {
		panic(err)
	}
	return binary.LittleEndian.Uint64(b)
}
