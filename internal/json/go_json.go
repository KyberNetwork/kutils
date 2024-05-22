// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build go_json
// +build go_json

package json

import "github.com/goccy/go-json"

type (
	// Unmarshaler is exported by go-json package.
	Unmarshaler = json.Unmarshaler
	// RawMessage is exported by go-json package.
	RawMessage = json.RawMessage
)

var (
	// Marshal is exported by go-json package.
	Marshal = json.Marshal
	// Unmarshal is exported by go-json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by go-json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by go-json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by go-json package.
	NewEncoder = json.NewEncoder
)
