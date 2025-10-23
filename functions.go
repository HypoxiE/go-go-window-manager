package main

import "strconv"

func HexToUint32(s string) uint32 {
	if s[0] == '#' {
		s = s[1:]
	}
	if len(s) != 8 {
		panic("invalid color length")
	}
	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	g, _ := strconv.ParseUint(s[2:4], 16, 8)
	b, _ := strconv.ParseUint(s[4:6], 16, 8)
	a, _ := strconv.ParseUint(s[6:8], 16, 8)

	return uint32(a)<<24 | uint32(r)<<16 | uint32(g)<<8 | uint32(b)
}
