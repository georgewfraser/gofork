// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmd/compile/i/amd64"
	"cmd/compile/i/arm"
	"cmd/compile/i/arm64"
	"cmd/compile/i/gc"
	"cmd/compile/i/mips"
	"cmd/compile/i/mips64"
	"cmd/compile/i/ppc64"
	"cmd/compile/i/s390x"
	"cmd/compile/i/wasm"
	"cmd/compile/i/x86"
	"cmd/internal/objabi"
	"fmt"
	"log"
	"os"
)

var archInits = map[string]func(*gc.Arch){
	"386":      x86.Init,
	"amd64":    amd64.Init,
	"amd64p32": amd64.Init,
	"arm":      arm.Init,
	"arm64":    arm64.Init,
	"mips":     mips.Init,
	"mipsle":   mips.Init,
	"mips64":   mips64.Init,
	"mips64le": mips64.Init,
	"ppc64":    ppc64.Init,
	"ppc64le":  ppc64.Init,
	"s390x":    s390x.Init,
	"wasm":     wasm.Init,
}

func main() {
	// disable timestamps for reproducible output
	log.SetFlags(0)
	log.SetPrefix("compile: ")

	archInit, ok := archInits[objabi.GOARCH]
	if !ok {
		fmt.Fprintf(os.Stderr, "compile: unknown architecture %q\n", objabi.GOARCH)
		os.Exit(2)
	}

	gc.Main(archInit)
	gc.Exit(0)
}
