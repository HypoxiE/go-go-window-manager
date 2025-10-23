package main

import "github.com/jezek/xgb/xproto"

type Hotkey struct {
	Keys      []xproto.Keycode
	Action    func(Args)
	Arguments Args

	KeyMask BitFlags
}

func (a *Hotkey) InitMask() {
	a.KeyMask = ToBitFlags(a.Keys)
}

type Args struct {
	Strings  []string
	Integers []int64
	Floats   []float64
	Booleans []bool
}

const flags_size = 4

type BitFlags [flags_size]uint64 // 4 * 64 = 256
func (a BitFlags) And(b BitFlags) BitFlags {
	var r BitFlags
	for i := 0; i < flags_size; i++ {
		r[i] = a[i] & b[i]
	}
	return r
}
func (a BitFlags) Or(b BitFlags) BitFlags {
	var r BitFlags
	for i := 0; i < flags_size; i++ {
		r[i] = a[i] | b[i]
	}
	return r
}
func (a BitFlags) Xor(b BitFlags) BitFlags {
	var r BitFlags
	for i := 0; i < flags_size; i++ {
		r[i] = a[i] ^ b[i]
	}
	return r
}
func (a BitFlags) Equal(b BitFlags) bool {
	for i := 0; i < flags_size; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func (b *BitFlags) Set(bit xproto.Keycode) {
	b[bit/64] |= 1 << (bit % 64)
}
func (b *BitFlags) Clear(bit xproto.Keycode) {
	b[bit/64] &^= 1 << (bit % 64)
}
func (b *BitFlags) Toggle(bit xproto.Keycode) {
	b[bit/64] ^= 1 << (bit % 64)
}
func (b BitFlags) Get(bit xproto.Keycode) bool {
	return b[bit/64]&(1<<(bit%64)) != 0
}
func ToBitFlags(keycodes []xproto.Keycode) BitFlags {
	var result BitFlags
	for _, v := range keycodes {
		result.Set(v)
	}
	return result
}
