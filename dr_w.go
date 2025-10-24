package main

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var BarWindow xproto.Window
var XGBConn *xgb.Conn

func DrawBar() {
	setup := xproto.Setup(XGBConn)
	screen := setup.DefaultScreen(XGBConn)

	BarWindow, _ = xproto.NewWindowId(XGBConn)
	width := uint16(screen.WidthInPixels)
	height := uint16(BarHeight)

	xproto.CreateWindow(XGBConn, screen.RootDepth, BarWindow, screen.Root,
		0, 0, width, height, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{
			BarBackground,
			xproto.EventMaskExposure |
				xproto.EventMaskButtonPress |
				xproto.EventMaskPointerMotion |
				xproto.EventMaskKeyPress |
				xproto.EventMaskKeyRelease |
				xproto.EventMaskEnterWindow |
				xproto.EventMaskLeaveWindow,
		})

	BarVisible = !BarVisible
	ToggleBar(Args{})
}

func DrawCursor() {
	setup := xproto.Setup(XGBConn)
	screen := setup.DefaultScreen(XGBConn)

	fontId, _ := xproto.NewFontId(XGBConn)
	xproto.OpenFont(XGBConn, fontId, uint16(len("cursor")), "cursor")
	cursor, _ := xproto.NewCursorId(XGBConn)
	xproto.CreateGlyphCursor(XGBConn, cursor,
		fontId, fontId,
		68, 69,
		0x0000, 0x0000, 0x0000,
		0xffff, 0xffff, 0xffff)

	xproto.ChangeWindowAttributes(XGBConn, screen.Root, xproto.CwCursor, []uint32{uint32(cursor)})
}
