// Copyright 2017, Project ArteMisc
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*

 */
package auth // import "github.com/phasehq/libgodium/auth"

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"github.com/phasehq/libgodium"
)

const (
	HmacSha256_Bytes    = 32
	HmacSha256_KeyBytes = 32

	HmacSha512_Bytes    = 64
	HmacSha512_KeyBytes = 32

	HmacSha512256_Bytes    = 32
	HmacSha512256_KeyBytes = 32

	Primitive = "hmacsha512256"
	Bytes     = HmacSha512256_Bytes
	KeyBytes  = HmacSha512256_KeyBytes
)

// hmacImpl implements the godium.Auth API on top of golang's own hash and hmac
// implementations.
type hmacImpl struct {
	hash.Hash
	key []byte
}

// New
func New(key []byte) (auth godium.Auth) {
	auth = NewHmacSha512256(key)
	return
}

// NewHmacSha256
func NewHmacSha256(key []byte) (auth godium.Auth) {
	auth = &hmacImpl{
		Hash: hmac.New(sha256.New, key),
		key:  key,
	}
	return
}

// NewHmacSha512
func NewHmacSha512(key []byte) (auth godium.Auth) {
	auth = &hmacImpl{
		Hash: hmac.New(sha512.New, key),
		key:  key,
	}
	return
}

// NewHmacSha512256
func NewHmacSha512256(key []byte) (auth godium.Auth) {
	auth = &hmacImpl{
		Hash: hmac.New(sha512.New512_256, key),
		key:  key,
	}
	return
}

// Wipe
func (h *hmacImpl) Wipe() {
	godium.Wipe(h.key)
}

// Verify
func (h *hmacImpl) Verify(tag []byte) (valid bool) {
	valid = hmac.Equal(h.Sum(nil), tag)
	return
}

func (h *hmacImpl) Bytes() int    { return h.Size() }
func (h *hmacImpl) KeyBytes() int { return len(h.key) }
