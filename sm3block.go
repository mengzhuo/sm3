// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sm3

import "encoding/binary"

const (
	t0  = 0x79cc4519 // 0 ≤ j ≤ 15
	t16 = 0x7a879d8a // 16 ≤ j ≤ 63
)

func block(dig *digest, p []byte) {

	var (
		w  [68]uint32
		wd [64]uint32 // w'

		ss1, ss2, tt1, tt2 uint32
	)

	h0, h1, h2, h3 := dig.h[0], dig.h[1], dig.h[2], dig.h[3]
	h4, h5, h6, h7 := dig.h[4], dig.h[5], dig.h[6], dig.h[7]

	for len(p) >= chunk {

		// expand data
		for j := 0; j < 16; j++ {
			w[j] = binary.BigEndian.Uint32(p[j*4 : (j+1)*4])
		}
		for j := 16; j < 68; j++ {
			w[j] = p1(w[j-16]^w[j-9]^(w[j-3]<<15|w[j-3]>>(32-15))) ^
				(w[j-13]<<7 | w[j-13]>>(32-7)) ^ w[j-6]
		}
		for j := 0; j < 64; j++ {
			wd[j] = w[j] ^ w[j+4]
		}

		a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7

		// compress function
		for j := uint(0); j < 16; j++ {
			ss1 = (a<<12 | a>>(32-12)) + e + (t0<<j | t0>>(32-j))
			ss1 = ss1<<7 | ss1>>(32-7)

			ss2 = ss1 ^ (a<<12 | a>>(32-12))
			tt1 = ff0(a, b, c) + d + ss2 + wd[j]
			tt2 = gg0(e, f, g) + h + ss1 + w[j]
			d = c
			c = b<<9 | b>>(32-9)
			b = a
			a = tt1
			h = g
			g = f<<19 | f>>(32-19)
			f = e
			e = p0(tt2)
		}
		for j := uint(16); j < 64; j++ {
			ss1 = (a<<12 | a>>(32-12)) + e + (t16<<(j%32) | t16>>((32-j)%32))
			ss1 = ss1<<7 | ss1>>(32-7)

			ss2 = ss1 ^ (a<<12 | a>>(32-12))
			tt1 = ff16(a, b, c) + d + ss2 + wd[j]
			tt2 = gg16(e, f, g) + h + ss1 + w[j]
			d = c
			c = b<<9 | b>>(32-9)
			b = a
			a = tt1
			h = g
			g = f<<19 | f>>(32-19)
			f = e
			e = p0(tt2)
		}

		h0 ^= a
		h1 ^= b
		h2 ^= c
		h3 ^= d
		h4 ^= e
		h5 ^= f
		h6 ^= g
		h7 ^= h

		p = p[chunk:]
	}

	dig.h[0], dig.h[1], dig.h[2], dig.h[3] = h0, h1, h2, h3
	dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h4, h5, h6, h7
}

func ff0(x, y, z uint32) uint32 { return x ^ y ^ z }

func ff16(x, y, z uint32) uint32 { return (x & y) | (x & z) | (y & z) }

func gg0(x, y, z uint32) uint32 { return x ^ y ^ z }

func gg16(x, y, z uint32) uint32 { return (x & y) | (^x & z) }

func p0(x uint32) uint32 { return x ^ ((x << 9) | (x >> (32 - 9))) ^ ((x << 17) | (x >> (32 - 17))) }

func p1(x uint32) uint32 { return x ^ ((x << 15) | (x >> (32 - 15))) ^ ((x << 23) | (x >> (32 - 23))) }
