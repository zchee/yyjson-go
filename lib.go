// Copyright 2022 The yyjson-go Authors
// SPDX-License-Identifier: BSD-3-Clause

package yyjson

import (
	"fmt"

	"modernc.org/libc"
)

func cString(s string) uintptr {
	p, err := libc.CString(s)
	if err != nil {
		panic(fmt.Errorf("allocate CString: %w", err))
	}

	return p
}
