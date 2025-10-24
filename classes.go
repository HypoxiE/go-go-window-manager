package main

import "github.com/jezek/xgb/xproto"

type Hotkey struct {
	Modifier  uint16
	Key       xproto.Keycode
	Action    func(Args)
	Arguments Args
}
type Args struct {
	Strings  []string
	Integers []int64
	Floats   []float64
	Booleans []bool
}
