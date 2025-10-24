package main

import (
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

var XGBConn *xgbutil.XUtil

func main() {
	var err error

	XGBConn, err = xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	defer XGBConn.Conn().Close()

	keybind.Initialize(XGBConn)

	//setup := xproto.Setup(XGBConn)
	//screen := setup.DefaultScreen(XGBConn)
	//root := screen.Root

	//DrawBar()
	//DrawCursor()
	//DrawTags(3)

	for _, hk := range Hotkeys {
		mods, keys, err := keybind.ParseString(XGBConn, hk.Hotkey)
		if err != nil {
			log.Fatal(err)
		}

		for _, key := range keys {
			keybind.Grab(XGBConn, XGBConn.RootWin(), mods, key)

			xevent.KeyPressFun(func(X *xgbutil.XUtil, ev xevent.KeyPressEvent) {
				if ev.Detail == key && ev.State&mods == mods {
					hk.Action(hk.Arguments)
				}
			}).Connect(XGBConn, XGBConn.RootWin())
		}
	}

	xevent.Main(XGBConn)
}
