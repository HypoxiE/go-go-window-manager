package main

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func window() {
	conn, _ := xgb.NewConn()
	setup := xproto.Setup(conn)
	screen := setup.DefaultScreen(conn)

	// создаем ID окна
	win, _ := xproto.NewWindowId(conn)

	// создаем окно
	xproto.CreateWindow(
		conn,
		screen.RootDepth, // глубина цвета
		win,
		screen.Root,
		200, 200, // позиция
		400, 300, // ширина/высота
		0, // границы
		xproto.WindowClassInputOutput,
		screen.RootVisual,
		xproto.CwBackPixel|xproto.CwEventMask,
		[]uint32{screen.WhitePixel, xproto.EventMaskExposure | xproto.EventMaskKeyPress},
	)

	// отображаем окно
	xproto.MapWindow(conn, win)
	//conn.Flush()

	// цикл событий
	for {
		ev, err := conn.WaitForEvent()
		if err != nil {
			panic(err)
		}

		switch e := ev.(type) {
		case xproto.ExposeEvent:
			println("Окно отрисовано!")
		case xproto.KeyPressEvent:
			println("Нажата клавиша — выходим.")
			return
		default:
			_ = e
		}
	}
}
