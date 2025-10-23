package main

import (
	"log"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var RuntimeFlag = true
var CurrentKeyMask BitFlags

func main() {

	for i := 0; i < len(Hotkeys); i++ {
		Hotkeys[i].InitMask()
	}

	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)
	root := screen.Root

	win, _ := xproto.NewWindowId(conn)
	width := uint16(screen.WidthInPixels)
	height := uint16(screen.HeightInPixels / 20)

	xproto.CreateWindow(conn, screen.RootDepth, win, root,
		0, 0, width, height, 0,
		xproto.WindowClassInputOutput, screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{BarColor, xproto.EventMaskExposure})

	//xproto.ChangeWindowAttributes(conn, root, xproto.CwEventMask,
	//	[]uint32{xproto.EventMaskKeyPress | xproto.EventMaskKeyRelease})

	//xproto.MapWindow(conn, win)

	for RuntimeFlag {
		ev, err := conn.WaitForEvent()
		if err != nil {
			log.Fatal(err)
		}

		switch e := ev.(type) {
		case xproto.ExposeEvent:
			gc, _ := xproto.NewGcontextId(conn)
			xproto.CreateGC(conn, gc, xproto.Drawable(win), xproto.GcForeground, []uint32{BarColor})
			xproto.PolyRectangle(conn, xproto.Drawable(win), gc, []xproto.Rectangle{
				{X: 0, Y: 0, Width: width, Height: uint16(BarHeight)},
			})

		case xproto.KeyPressEvent:
			//fmt.Printf("Key pressed: keycode=%d\n", e.Detail)
			// 00000001|10000000

			CurrentKeyMask.Set(e.Detail)

			for _, hk := range Hotkeys {
				if CurrentKeyMask.And(hk.KeyMask) == hk.KeyMask {
					hk.Action(hk.Arguments)
				}
			}

			//for _, hk := range Hotkeys {
			//	if hk.Key == 0 && e.State&hk.Modifiers == hk.Modifiers {
			//		hk.Action(hk.Arguments)
			//	} else if e.Detail == xproto.Keycode(hk.Key) && e.State&hk.Modifiers == hk.Modifiers {
			//		hk.Action(hk.Arguments)
			//	}
			//}
		case xproto.KeyReleaseEvent:
			//fmt.Printf("Key released: keycode=%d\n", e.Detail)
			CurrentKeyMask.Clear(e.Detail)
		}
	}
}
