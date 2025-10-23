package main

import (
	"fwm/xkeys"

	"github.com/jezek/xgb/xproto"
)

const MODKEY = xkeys.XK_Super_L

// Interfase
var (
	BarVisible = true
)

const (
	BarHeight uint32 = 24
)

// Colors
var (
	BarBackground = HexToUint32("#444444ff")

	ActiveBorderColor   = HexToUint32("#0076ecff")
	DeactiveBorderColor = HexToUint32("#585858ff")
)

var Hotkeys = []Hotkey{
	{Keys: []xproto.Keycode{xkeys.XK_0}, Action: SayHello, Arguments: Args{}},
	{Keys: []xproto.Keycode{MODKEY, xkeys.XK_q}, Action: Quit, Arguments: Args{}},
	{Keys: []xproto.Keycode{xkeys.XK_q, xkeys.XK_w}, Action: ToggleBar, Arguments: Args{}},
	//{Key: 24, Modifiers: xproto.ModMask1, Action: func() { fmt.Println("Alt+Q") }},
}

var Tags = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
