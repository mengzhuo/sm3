sm3
=======

[![Build Status](https://travis-ci.org/mengzhuo/sm3.svg?branch=master)](https://travis-ci.org/mengzhuo/sm3)
[![GoDoc](https://godoc.org/github.com/mengzhuo/sm3?status.svg)](https://godoc.org/github.com/mengzhuo/sm3)

Go implement of sm3 hash algorithm

Example
--------

```
h := sm3.New()
data := "The answer to the ultimate question of life, the universe and everything is 42."
io.WriteString(h, data)
fmt.Printf("%x", h.Sum(nil))
```
