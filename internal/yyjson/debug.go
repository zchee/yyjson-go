// Copyright 2022 The yyjson-go Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build yyjson.debug

package yyjson

import (
	"github.com/Code-Hex/dd/p"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/styles"
)

// Dump dumps a for debug.
func Dump(a ...interface{}) {
	p.New(p.WithStyle(styles.DoomOne2), p.WithFormatter(formatters.TTY16m)).P(a...)
}
