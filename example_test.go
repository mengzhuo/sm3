// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm3_test

import (
	"fmt"
	"io"

	"github.com/mengzhuo/sm3"
)

func ExampleNew() {
	h := sm3.New()
	data := "The answer to the ultimate question of life, the universe and everything is 42."
	io.WriteString(h, data)
	fmt.Printf("%x", h.Sum(nil))
	// Output: 9236124594a6d02f0639c4e92916911076b98e21fc09b12078ba87eaadb3dea4
}

func ExampleSum() {
	data := "The answer to the ultimate question of life, the universe and everything is 42."
	fmt.Printf("%x", sm3.Sum([]byte(data)))
	// Output: 9236124594a6d02f0639c4e92916911076b98e21fc09b12078ba87eaadb3dea4
}
