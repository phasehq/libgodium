// Copyright 2017, Project ArteMisc
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*

Package Godium provides implementations for the primitives present in the
Libsodium library.

The library is fully written in Go (or go-assembly), and based on interfaces
found in Go's standard library.

*/
package godium // import "github.com/phasehq/libgodium"

// Version information, represents the latest libsodium version that this build
// is compatible with.
const (
	VersionMajor = 10
	VersionMinor = 0

	Version = "1.0.15"
)
