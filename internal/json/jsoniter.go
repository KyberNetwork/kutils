// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build jsoniter
// +build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

type (
	// Unmarshaler is exported by jsoniter package.
	Unmarshaler = json.Unmarshaler
	// RawMessage is exported by jsoniter package.
	RawMessage = json.RawMessage
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by jsoniter package.
	Marshal = json.Marshal
	// Unmarshal is exported by jsoniter package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by jsoniter package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by jsoniter package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by jsoniter package.
	NewEncoder = json.NewEncoder
)
