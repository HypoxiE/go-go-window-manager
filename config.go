package main

import (
	"fwm/xkeys"

	"github.com/jezek/xgb/xproto"
)

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

const MODMASK = xproto.ModMask1

/*
 */
var Hotkeys = []Hotkey{
	{Modifier: MODMASK, Key: xkeys.XK_0, Action: SayHello, Arguments: Args{}},
	{Modifier: MODMASK, Key: xkeys.XK_q, Action: Quit, Arguments: Args{}},
	{Modifier: MODMASK, Key: xkeys.XK_1, Action: ToggleBar, Arguments: Args{}},
	//{Key: 24, Modifiers: xproto.ModMask1, Action: func() { fmt.Println("Alt+Q") }},
}

var Tags = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
