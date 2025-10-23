package xkeys

import "github.com/jezek/xgb/xproto"

// Keyboard key constants (X11 keycodes)
const (
	// Modifier keys
	XK_Escape           xproto.Keycode = 9
	XK_Tab              xproto.Keycode = 23
	XK_Return           xproto.Keycode = 36
	XK_BackSpace        xproto.Keycode = 22
	XK_Control_L        xproto.Keycode = 37
	XK_Control_R        xproto.Keycode = 105
	XK_Shift_L          xproto.Keycode = 50
	XK_Shift_R          xproto.Keycode = 62
	XK_Alt_L            xproto.Keycode = 64
	XK_Alt_R            xproto.Keycode = 108
	XK_Super_L          xproto.Keycode = 133
	XK_Super_R          xproto.Keycode = 134
	XK_Menu             xproto.Keycode = 135
	XK_ISO_Level3_Shift xproto.Keycode = 92
	XK_ISO_Level5_Shift xproto.Keycode = 203

	// Alphabet keys
	XK_a xproto.Keycode = 38
	XK_b xproto.Keycode = 56
	XK_c xproto.Keycode = 54
	XK_d xproto.Keycode = 40
	XK_e xproto.Keycode = 26
	XK_f xproto.Keycode = 41
	XK_g xproto.Keycode = 42
	XK_h xproto.Keycode = 43
	XK_i xproto.Keycode = 31
	XK_j xproto.Keycode = 44
	XK_k xproto.Keycode = 45
	XK_l xproto.Keycode = 46
	XK_m xproto.Keycode = 58
	XK_n xproto.Keycode = 57
	XK_o xproto.Keycode = 32
	XK_p xproto.Keycode = 33
	XK_q xproto.Keycode = 24
	XK_r xproto.Keycode = 27
	XK_s xproto.Keycode = 39
	XK_t xproto.Keycode = 28
	XK_u xproto.Keycode = 30
	XK_v xproto.Keycode = 55
	XK_w xproto.Keycode = 25
	XK_x xproto.Keycode = 53
	XK_y xproto.Keycode = 29
	XK_z xproto.Keycode = 52

	// Number keys
	XK_0 xproto.Keycode = 19
	XK_1 xproto.Keycode = 10
	XK_2 xproto.Keycode = 11
	XK_3 xproto.Keycode = 12
	XK_4 xproto.Keycode = 13
	XK_5 xproto.Keycode = 14
	XK_6 xproto.Keycode = 15
	XK_7 xproto.Keycode = 16
	XK_8 xproto.Keycode = 17
	XK_9 xproto.Keycode = 18

	// Symbols
	XK_minus        xproto.Keycode = 20
	XK_equal        xproto.Keycode = 21
	XK_bracketleft  xproto.Keycode = 34
	XK_bracketright xproto.Keycode = 35
	XK_semicolon    xproto.Keycode = 47
	XK_apostrophe   xproto.Keycode = 48
	XK_grave        xproto.Keycode = 49
	XK_backslash    xproto.Keycode = 51
	XK_comma        xproto.Keycode = 59
	XK_period       xproto.Keycode = 60
	XK_slash        xproto.Keycode = 61
	XK_space        xproto.Keycode = 65

	// Function keys
	XK_F1  xproto.Keycode = 67
	XK_F2  xproto.Keycode = 68
	XK_F3  xproto.Keycode = 69
	XK_F4  xproto.Keycode = 70
	XK_F5  xproto.Keycode = 71
	XK_F6  xproto.Keycode = 72
	XK_F7  xproto.Keycode = 73
	XK_F8  xproto.Keycode = 74
	XK_F9  xproto.Keycode = 75
	XK_F10 xproto.Keycode = 76
	XK_F11 xproto.Keycode = 95
	XK_F12 xproto.Keycode = 96

	// Navigation keys
	XK_Up     xproto.Keycode = 111
	XK_Down   xproto.Keycode = 116
	XK_Left   xproto.Keycode = 113
	XK_Right  xproto.Keycode = 114
	XK_Home   xproto.Keycode = 110
	XK_End    xproto.Keycode = 115
	XK_Prior  xproto.Keycode = 112
	XK_Next   xproto.Keycode = 117
	XK_Insert xproto.Keycode = 118
	XK_Delete xproto.Keycode = 119
)
