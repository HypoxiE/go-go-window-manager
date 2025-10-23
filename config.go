package main

import (
	"fwm/xkeys"

	"github.com/jezek/xgb/xproto"
)

const MODKEY = xkeys.XK_Super_L

// Interfase
var (
	BarHeight = 200
)

// Colors
var (
	BarColor = HexToUint32("#444444ff")

	ActiveBorderColor   = HexToUint32("#0076ecff")
	DeactiveBorderColor = HexToUint32("#585858ff")
)

var Hotkeys = []Hotkey{
	{Keys: []xproto.Keycode{xkeys.XK_0}, Action: SayHello, Arguments: Args{}},
	{Keys: []xproto.Keycode{MODKEY, xkeys.XK_q}, Action: Quit, Arguments: Args{}},
	//{Key: 24, Modifiers: xproto.ModMask1, Action: func() { fmt.Println("Alt+Q") }},
}
