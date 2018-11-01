// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elib

import (
	"fmt"
)

type Logger interface {
	Logf(format string, args ...interface{})
	Logln(args ...interface{})
}

var DbgElib bool = false

func DbgElibLog(argfmt string, arg ...interface{}) {
	if DbgElib {
		fmt.Printf(argfmt, arg...)
	}
}

func DbgElibLogEnable(f bool) {
	DbgElib = f
}
