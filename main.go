// Copyright (c) 2014, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	conf, cmd, err := NewConfigFromFlags()
	if err != nil {
		fmt.Println("Error:", err, "\n")
		Usage(nil)
		os.Exit(1)
	}

	cmd(conf)
}