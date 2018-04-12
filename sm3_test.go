// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm3

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"testing"
)

type sm3Test struct {
	in  string
	out string
	hex bool
}

var golden = []sm3Test{
	{"", "1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b", false},
	// example 1
	{"abc", "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0", false},
	// example 2
	{strings.Repeat("abcd", 16), "debe9ff92275b8a138604889c18e5a4d6fdb70e5387e5765293dcba39c0c5732", false},
	{"ac1fb06e4dee27f7", "51f8e3f42d026d7a6876969b00281c9d5f62374936c88b78043b21be851fed3e", true},
	{"c4a65db51e1f78a2", "4b314697e1cacd324dc39d8bb9ca6d7a68918f94f7e232ae987e346b81d9de92", true},
	{"798cbf99a4e3fe3e", "87569c304bea43e06d53235a2e0b3e5475c1bd2cc3d7ed6d7a68be0bf1df51fc", true},
	{"98ab873d01bc4b60d4cf951ff33a21f2", "d1cb66c3e64b493112da1bce4a5ecd82aba4c76d7a6813dd59c4a3850a5952be", true},
	{"ad42e9cfe799a66c514477b10e8d5228", "0aa079f40c13fb39f21d9c8b9c1d0e2841bf4235316d7a68b96d10cb95e613eb", true},
	{"7c74e3b099bdd09e502e3654474ef185", "f87330d79aa0a609d2bdeb0a54d77ace85892c70618fed432ffa3bf8316d7a68", true},
}

func TestGolden(t *testing.T) {

	for i := 0; i < len(golden); i++ {
		g := golden[i]
		if g.hex {
			x, err := hex.DecodeString(g.in)
			if err != nil {
				t.Fatalf("decode failed:%s %s", g.in, err)
			}
			g.in = string(x)

		}
		s := fmt.Sprintf("%x", Sum([]byte(g.in)))
		if s != g.out {
			t.Fatalf("Sum function: sm3(%s) \ngot : %s \nwant: %s",
				g.in, s, g.out)
		}

		c := New()
		for j := 0; j < 2; j++ {
			var sum []byte
			switch j {
			case 0:
				io.WriteString(c, g.in)
				sum = c.Sum(nil)
			case 1:
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
				sum = c.Sum(nil)
			}
			s := fmt.Sprintf("%x", sum)
			if s != g.out {
				t.Fatalf("Sum function: sm3(%s)[%d] \ngot : %s \nwant: %s",
					g.in, j, s, g.out)
			}
			c.Reset()
		}
	}
}

var bench = New()
var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash(b *testing.B) {
	benchSize := []int{8, 320, 1024, 8192}
	rand.Read(buf)
	b.ResetTimer()
	for _, size := range benchSize {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			benchmarkSize(b, size)
		})
	}
}
