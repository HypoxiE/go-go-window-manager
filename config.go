package main

// Interfase
var (
	BarVisible = true
)

const (
	BarHeight uint32 = 24
)

// Colors
var (
	BarBackgroundColor = HexToUint32("#444444ff")

	TagsActiveBackgroundColor    = HexToUint32("#758beeff")
	TagsNotActiveBackgroundColor = HexToUint32("#943838ff")

	TagsActiveTextColor    = HexToUint32("#000000ff")
	TagsNotActiveTextColor = HexToUint32("#000000ff")

	ActiveBorderColor    = HexToUint32("#0076ecff")
	NotActiveBorderColor = HexToUint32("#585858ff")
)

/*
 */
var Hotkeys = []Hotkey{
	{Hotkey: "mod1-0", Action: SayHello, Arguments: Args{Strings: []string{"world"}}},
	{Hotkey: "mod1-q", Action: Quit, Arguments: Args{}},
	//{Modifier: MODMASK, Key: "1", Action: ToggleBar, Arguments: Args{}},
	{Hotkey: "7", Action: SayHello, Arguments: Args{}},
	//{Key: 24, Modifiers: xproto.ModMask1, Action: func() { fmt.Println("Alt+Q") }},
}

var Tags = []string{"a", "b", "3", "4", "5", "6", "7", "8", "9"}
