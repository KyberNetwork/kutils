// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter && !go_json && !(sonic && avx && (linux || windows || darwin) && amd64)
// +build !jsoniter
// +build !go_json
// +build !sonic !avx !linux,!windows,!darwin !amd64

package json

import "encoding/json"

type (
	// Unmarshaler is exported by std json package.
	Unmarshaler = json.Unmarshaler
	// RawMessage is exported by std json package.
	RawMessage = json.RawMessage
)

var (
	// Marshal is exported by std json package.
	Marshal = json.Marshal
	// Unmarshal is exported by std json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by std json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by std json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by std json package.
	NewEncoder = json.NewEncoder
)
