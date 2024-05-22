// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build sonic && avx && (linux || windows || darwin) && amd64
// +build sonic
// +build avx
// +build linux windows darwin
// +build amd64

package json

import "github.com/bytedance/sonic"

type (
	// Unmarshaler is exported by sonic package.
	Unmarshaler = json.Unmarshaler
	// RawMessage is exported by sonic package.
	RawMessage = json.RawMessage
)

var (
	json = sonic.ConfigStd
	// Marshal is exported by sonic package.
	Marshal = json.Marshal
	// Unmarshal is exported by sonic package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by sonic package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by sonic package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by sonic package.
	NewEncoder = json.NewEncoder
)
