// Copyright 2017, Project ArteMisc
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*

 */
package hash // import "github.com/phasehq/libgodium/hash"

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/phasehq/libgodium"
)

const (
	Sha256_Bytes = 32

	Sha512_Bytes = 64

	Primitive = "sha512"
	Bytes     = Sha512_Bytes
)

// shaImpl
type shaImpl struct {
	hash.Hash
}

// New
func New() (h godium.Hash) {
	h = NewSha512()
	return
}

// NewSha256
func NewSha256() (h godium.Hash) {
	h = &shaImpl{
		Hash: sha256.New(),
	}
	return
}

// NewSha256
func NewSha512() (h godium.Hash) {
	h = &shaImpl{
		Hash: sha512.New(),
	}
	return
}

// Sum
func Sum(dst, data []byte) (sum []byte) {
	sum = SumSha512(dst, data)
	return
}

// SumSha256
func SumSha256(dst, data []byte) (sum []byte) {
	sha := sha256.Sum256(data)
	sum = append(dst, sha[:]...)
	return
}

// SumSha512
func SumSha512(dst, data []byte) (sum []byte) {
	sha := sha512.Sum512(data)
	sum = append(dst, sha[:]...)
	return
}

func (s *shaImpl) Bytes() int { return s.Hash.Size() }
