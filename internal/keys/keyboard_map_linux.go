package keys

import "github.com/micmonay/keybd_event"

var keyboardMap = map[string]int{
	/* SHARED KEYS */
	"VK_SP1":       keybd_event.VK_SP1,
	"VK_SP2":       keybd_event.VK_SP2,
	"VK_SP3":       keybd_event.VK_SP3,
	"VK_SP4":       keybd_event.VK_SP4,
	"VK_SP5":       keybd_event.VK_SP5,
	"VK_SP6":       keybd_event.VK_SP6,
	"VK_SP7":       keybd_event.VK_SP7,
	"VK_SP8":       keybd_event.VK_SP8,
	"VK_SP9":       keybd_event.VK_SP9,
	"VK_SP10":      keybd_event.VK_SP10,
	"VK_SP11":      keybd_event.VK_SP11,
	"VK_SP12":      keybd_event.VK_SP12,
	"VK_ESC":       keybd_event.VK_ESC,
	"VK_1":         keybd_event.VK_1,
	"VK_2":         keybd_event.VK_2,
	"VK_3":         keybd_event.VK_3,
	"VK_4":         keybd_event.VK_4,
	"VK_5":         keybd_event.VK_5,
	"VK_6":         keybd_event.VK_6,
	"VK_7":         keybd_event.VK_7,
	"VK_8":         keybd_event.VK_8,
	"VK_9":         keybd_event.VK_9,
	"VK_0":         keybd_event.VK_0,
	"VK_Q":         keybd_event.VK_Q,
	"VK_W":         keybd_event.VK_W,
	"VK_E":         keybd_event.VK_E,
	"VK_R":         keybd_event.VK_R,
	"VK_T":         keybd_event.VK_T,
	"VK_Y":         keybd_event.VK_Y,
	"VK_U":         keybd_event.VK_U,
	"VK_I":         keybd_event.VK_I,
	"VK_O":         keybd_event.VK_O,
	"VK_P":         keybd_event.VK_P,
	"VK_A":         keybd_event.VK_A,
	"VK_S":         keybd_event.VK_S,
	"VK_D":         keybd_event.VK_D,
	"VK_F":         keybd_event.VK_F,
	"VK_G":         keybd_event.VK_G,
	"VK_H":         keybd_event.VK_H,
	"VK_J":         keybd_event.VK_J,
	"VK_K":         keybd_event.VK_K,
	"VK_L":         keybd_event.VK_L,
	"VK_Z":         keybd_event.VK_Z,
	"VK_X":         keybd_event.VK_X,
	"VK_C":         keybd_event.VK_C,
	"VK_V":         keybd_event.VK_V,
	"VK_B":         keybd_event.VK_B,
	"VK_N":         keybd_event.VK_N,
	"VK_M":         keybd_event.VK_M,
	"VK_F1":        keybd_event.VK_F1,
	"VK_F2":        keybd_event.VK_F2,
	"VK_F3":        keybd_event.VK_F3,
	"VK_F4":        keybd_event.VK_F4,
	"VK_F5":        keybd_event.VK_F5,
	"VK_F6":        keybd_event.VK_F6,
	"VK_F7":        keybd_event.VK_F7,
	"VK_F8":        keybd_event.VK_F8,
	"VK_F9":        keybd_event.VK_F9,
	"VK_F10":       keybd_event.VK_F10,
	"VK_F11":       keybd_event.VK_F11,
	"VK_F12":       keybd_event.VK_F12,
	"VK_F13":       keybd_event.VK_F13,
	"VK_F14":       keybd_event.VK_F14,
	"VK_F15":       keybd_event.VK_F15,
	"VK_F16":       keybd_event.VK_F16,
	"VK_F17":       keybd_event.VK_F17,
	"VK_F18":       keybd_event.VK_F18,
	"VK_F19":       keybd_event.VK_F19,
	"VK_F20":       keybd_event.VK_F20,
	"VK_MINUS":     keybd_event.VK_MINUS,
	"VK_EQUAL":     keybd_event.VK_EQUAL,
	"VK_TAB":       keybd_event.VK_TAB,
	"VK_ENTER":     keybd_event.VK_ENTER,
	"VK_SEMICOLON": keybd_event.VK_SEMICOLON,
	"VK_GRAVE":     keybd_event.VK_GRAVE,
	"VK_BACKSLASH": keybd_event.VK_BACKSLASH,
	"VK_COMMA":     keybd_event.VK_COMMA,
	"VK_SLASH":     keybd_event.VK_SLASH,
	"VK_SPACE":     keybd_event.VK_SPACE,
	"VK_CAPSLOCK":  keybd_event.VK_CAPSLOCK,
	"VK_PAGEUP":    keybd_event.VK_PAGEUP,
	"VK_PAGEDOWN":  keybd_event.VK_PAGEDOWN,
	"VK_END":       keybd_event.VK_END,
	"VK_HOME":      keybd_event.VK_HOME,
	"VK_LEFT":      keybd_event.VK_LEFT,
	"VK_UP":        keybd_event.VK_UP,
	"VK_RIGHT":     keybd_event.VK_RIGHT,
	"VK_DOWN":      keybd_event.VK_DOWN,
	"VK_DELETE":    keybd_event.VK_DELETE,
	"VK_HELP":      keybd_event.VK_HELP, /* Linux: AL Integrated Help Center;  */

	/* WINDOWS AND LINUX KEYS */
	"VK_F21":        keybd_event.VK_F21,
	"VK_F22":        keybd_event.VK_F22,
	"VK_F23":        keybd_event.VK_F23,
	"VK_F24":        keybd_event.VK_F24,
	"VK_NUMLOCK":    keybd_event.VK_NUMLOCK,
	"VK_SCROLLLOCK": keybd_event.VK_SCROLLLOCK,
	"VK_RESERVED":   keybd_event.VK_RESERVED,
	"VK_BACKSPACE":  keybd_event.VK_BACKSPACE,
	"VK_LEFTBRACE":  keybd_event.VK_LEFTBRACE,
	"VK_RIGHTBRACE": keybd_event.VK_RIGHTBRACE,
	"VK_APOSTROPHE": keybd_event.VK_APOSTROPHE,
	"VK_DOT":        keybd_event.VK_DOT,
	"VK_KPASTERISK": keybd_event.VK_KPASTERISK,
	"VK_KP0":        keybd_event.VK_KP0,
	"VK_KP1":        keybd_event.VK_KP1,
	"VK_KP2":        keybd_event.VK_KP2,
	"VK_KP3":        keybd_event.VK_KP3,
	"VK_KP4":        keybd_event.VK_KP4,
	"VK_KP5":        keybd_event.VK_KP5,
	"VK_KP6":        keybd_event.VK_KP6,
	"VK_KP7":        keybd_event.VK_KP7,
	"VK_KP8":        keybd_event.VK_KP8,
	"VK_KP9":        keybd_event.VK_KP9,
	"VK_KPMINUS":    keybd_event.VK_KPMINUS,
	"VK_KPPLUS":     keybd_event.VK_KPPLUS,
	"VK_KPDOT":      keybd_event.VK_KPDOT,
	"VK_CANCEL":     keybd_event.VK_CANCEL, /* Linux: AC Cancel;  */
	"VK_BACK":       keybd_event.VK_BACK,   /* Linux: AC Back;  */
	"VK_PAUSE":      keybd_event.VK_PAUSE,
	"VK_HANGUEL":    keybd_event.VK_HANGUEL,
	"VK_HANJA":      keybd_event.VK_HANJA,
	"VK_PRINT":      keybd_event.VK_PRINT, /* Linux: AC Print;  */
	"VK_INSERT":     keybd_event.VK_INSERT,
	"VK_PLAY":       keybd_event.VK_PLAY,

	/* LINUX AND MAC KEYS */
	"VK_MUTE":       keybd_event.VK_MUTE,
	"VK_VOLUMEDOWN": keybd_event.VK_VOLUMEDOWN,
	"VK_VOLUMEUP":   keybd_event.VK_VOLUMEUP,

	/* ONLY LINUX KEYS */
	"VK_KPJPCOMMA":        keybd_event.VK_KPJPCOMMA,
	"VK_KPENTER":          keybd_event.VK_KPENTER,
	"VK_KPSLASH":          keybd_event.VK_KPSLASH,
	"VK_KPEQUAL":          keybd_event.VK_KPEQUAL,
	"VK_KPPLUSMINUS":      keybd_event.VK_KPPLUSMINUS,
	"VK_KPCOMMA":          keybd_event.VK_KPCOMMA,
	"VK_ZENKAKUHANKAKU":   keybd_event.VK_ZENKAKUHANKAKU,
	"VK_102ND":            keybd_event.VK_102ND,
	"VK_RO":               keybd_event.VK_RO,
	"VK_KATAKANA":         keybd_event.VK_KATAKANA,
	"VK_HIRAGANA":         keybd_event.VK_HIRAGANA,
	"VK_HENKAN":           keybd_event.VK_HENKAN,
	"VK_KATAKANAHIRAGANA": keybd_event.VK_KATAKANAHIRAGANA,
	"VK_MUHENKAN":         keybd_event.VK_MUHENKAN,
	"VK_SYSRQ":            keybd_event.VK_SYSRQ,
	"VK_LINEFEED":         keybd_event.VK_LINEFEED,
	"VK_MACRO":            keybd_event.VK_MACRO,
	"VK_POWER":            keybd_event.VK_POWER, /* Linux: SC System Power Down;  */
	"VK_SCALE":            keybd_event.VK_SCALE, /* Linux: AL Compiz Scale (Expose);  */
	"VK_HANGEUL":          keybd_event.VK_HANGEUL,
	"VK_YEN":              keybd_event.VK_YEN,
	"VK_LEFTMETA":         keybd_event.VK_LEFTMETA,
	"VK_RIGHTMETA":        keybd_event.VK_RIGHTMETA,
	"VK_COMPOSE":          keybd_event.VK_COMPOSE,
	"VK_STOP":             keybd_event.VK_STOP, /* Linux: AC Stop;  */
	"VK_AGAIN":            keybd_event.VK_AGAIN,
	"VK_PROPS":            keybd_event.VK_PROPS, /* Linux: AC Properties;  */
	"VK_UNDO":             keybd_event.VK_UNDO,  /* Linux: AC Undo;  */
	"VK_FRONT":            keybd_event.VK_FRONT,
	"VK_COPY":             keybd_event.VK_COPY,  /* Linux: AC Copy;  */
	"VK_OPEN":             keybd_event.VK_OPEN,  /* Linux: AC Open;  */
	"VK_PASTE":            keybd_event.VK_PASTE, /* Linux: AC Paste;  */
	"VK_FIND":             keybd_event.VK_FIND,  /* Linux: AC Search;  */
	"VK_CUT":              keybd_event.VK_CUT,   /* Linux: AC Cut;  */
	"VK_MENU":             keybd_event.VK_MENU,  /* Linux: Menu (show menu);  */
	"VK_CALC":             keybd_event.VK_CALC,  /* Linux: AL Calculator;  */
	"VK_SETUP":            keybd_event.VK_SETUP,
	"VK_SLEEP":            keybd_event.VK_SLEEP,  /* Linux: SC System Sleep;  */
	"VK_WAKEUP":           keybd_event.VK_WAKEUP, /* Linux: System Wake Up;  */
	"VK_FILE":             keybd_event.VK_FILE,   /* Linux: AL Local Machine Browser;  */
	"VK_SENDFILE":         keybd_event.VK_SENDFILE,
	"VK_DELETEFILE":       keybd_event.VK_DELETEFILE,
	"VK_XFER":             keybd_event.VK_XFER,
	"VK_PROG1":            keybd_event.VK_PROG1,
	"VK_PROG2":            keybd_event.VK_PROG2,
	"VK_WWW":              keybd_event.VK_WWW, /* Linux: AL Internet Browser;  */
	"VK_MSDOS":            keybd_event.VK_MSDOS,
	"VK_COFFEE":           keybd_event.VK_COFFEE, /* Linux: AL Terminal Lock/Screensaver;  */
	"VK_SCREENLOCK":       keybd_event.VK_SCREENLOCK,
	"VK_ROTATE_DISPLAY":   keybd_event.VK_ROTATE_DISPLAY, /* Linux: Display orientation for e.g. tablets;  */
	"VK_DIRECTION":        keybd_event.VK_DIRECTION,
	"VK_CYCLEWINDOWS":     keybd_event.VK_CYCLEWINDOWS,
	"VK_MAIL":             keybd_event.VK_MAIL,
	"VK_BOOKMARKS":        keybd_event.VK_BOOKMARKS, /* Linux: AC Bookmarks;  */
	"VK_COMPUTER":         keybd_event.VK_COMPUTER,
	"VK_FORWARD":          keybd_event.VK_FORWARD, /* Linux: AC Forward;  */
	"VK_CLOSECD":          keybd_event.VK_CLOSECD,
	"VK_EJECTCD":          keybd_event.VK_EJECTCD,
	"VK_EJECTCLOSECD":     keybd_event.VK_EJECTCLOSECD,
	"VK_NEXTSONG":         keybd_event.VK_NEXTSONG,
	"VK_PLAYPAUSE":        keybd_event.VK_PLAYPAUSE,
	"VK_PREVIOUSSONG":     keybd_event.VK_PREVIOUSSONG,
	"VK_STOPCD":           keybd_event.VK_STOPCD,
	"VK_RECORD":           keybd_event.VK_RECORD,
	"VK_REWIND":           keybd_event.VK_REWIND,
	"VK_PHONE":            keybd_event.VK_PHONE, /* Linux: Media Select Telephone;  */
	"VK_ISO":              keybd_event.VK_ISO,
	"VK_CONFIG":           keybd_event.VK_CONFIG,   /* Linux: AL Consumer Control Configuration;  */
	"VK_HOMEPAGE":         keybd_event.VK_HOMEPAGE, /* Linux: AC Home;  */
	"VK_REFRESH":          keybd_event.VK_REFRESH,  /* Linux: AC Refresh;  */
	"VK_EXIT":             keybd_event.VK_EXIT,     /* Linux: AC Exit;  */
	"VK_MOVE":             keybd_event.VK_MOVE,
	"VK_EDIT":             keybd_event.VK_EDIT,
	"VK_SCROLLUP":         keybd_event.VK_SCROLLUP,
	"VK_SCROLLDOWN":       keybd_event.VK_SCROLLDOWN,
	"VK_KPLEFTPAREN":      keybd_event.VK_KPLEFTPAREN,
	"VK_KPRIGHTPAREN":     keybd_event.VK_KPRIGHTPAREN,
	"VK_NEW":              keybd_event.VK_NEW,  /* Linux: AC New;  */
	"VK_REDO":             keybd_event.VK_REDO, /* Linux: AC Redo/Repeat;  */
	"VK_PLAYCD":           keybd_event.VK_PLAYCD,
	"VK_PAUSECD":          keybd_event.VK_PAUSECD,
	"VK_PROG3":            keybd_event.VK_PROG3,
	"VK_PROG4":            keybd_event.VK_PROG4,
	"VK_DASHBOARD":        keybd_event.VK_DASHBOARD, /* Linux: AL Dashboard;  */
	"VK_SUSPEND":          keybd_event.VK_SUSPEND,
	"VK_CLOSE":            keybd_event.VK_CLOSE, /* Linux: AC Close;  */
	"VK_FASTFORWARD":      keybd_event.VK_FASTFORWARD,
	"VK_BASSBOOST":        keybd_event.VK_BASSBOOST,
	"VK_HP":               keybd_event.VK_HP,
	"VK_CAMERA":           keybd_event.VK_CAMERA,
	"VK_SOUND":            keybd_event.VK_SOUND,
	"VK_QUESTION":         keybd_event.VK_QUESTION,
	"VK_EMAIL":            keybd_event.VK_EMAIL,
	"VK_CHAT":             keybd_event.VK_CHAT,
	"VK_SEARCH":           keybd_event.VK_SEARCH,
	"VK_CONNECT":          keybd_event.VK_CONNECT,
	"VK_FINANCE":          keybd_event.VK_FINANCE, /* Linux: AL Checkbook/Finance;  */
	"VK_SPORT":            keybd_event.VK_SPORT,
	"VK_SHOP":             keybd_event.VK_SHOP,
	"VK_ALTERASE":         keybd_event.VK_ALTERASE,
	"VK_BRIGHTNESSDOWN":   keybd_event.VK_BRIGHTNESSDOWN,
	"VK_BRIGHTNESSUP":     keybd_event.VK_BRIGHTNESSUP,
	"VK_MEDIA":            keybd_event.VK_MEDIA,
	"VK_SWITCHVIDEOMODE":  keybd_event.VK_SWITCHVIDEOMODE, /* Linux: Cycle between available video outputs (Monitor/LCD/TV-out/etc);  */
	"VK_KBDILLUMTOGGLE":   keybd_event.VK_KBDILLUMTOGGLE,
	"VK_KBDILLUMDOWN":     keybd_event.VK_KBDILLUMDOWN,
	"VK_KBDILLUMUP":       keybd_event.VK_KBDILLUMUP,
	"VK_SEND":             keybd_event.VK_SEND,        /* Linux: AC Send;  */
	"VK_REPLY":            keybd_event.VK_REPLY,       /* Linux: AC Reply;  */
	"VK_FORWARDMAIL":      keybd_event.VK_FORWARDMAIL, /* Linux: AC Forward Msg;  */
	"VK_SAVE":             keybd_event.VK_SAVE,        /* Linux: AC Save;  */
	"VK_DOCUMENTS":        keybd_event.VK_DOCUMENTS,
	"VK_BATTERY":          keybd_event.VK_BATTERY,
	"VK_BLUETOOTH":        keybd_event.VK_BLUETOOTH,
	"VK_WLAN":             keybd_event.VK_WLAN,
	"VK_UWB":              keybd_event.VK_UWB,
	"VK_UNKNOWN":          keybd_event.VK_UNKNOWN,
	"VK_VIDEO_NEXT":       keybd_event.VK_VIDEO_NEXT,       /* Linux: drive next video source;  */
	"VK_VIDEO_PREV":       keybd_event.VK_VIDEO_PREV,       /* Linux: drive previous video source;  */
	"VK_BRIGHTNESS_CYCLE": keybd_event.VK_BRIGHTNESS_CYCLE, /* Linux: brightness up, after max is min;  */
	"VK_BRIGHTNESS_AUTO":  keybd_event.VK_BRIGHTNESS_AUTO,  /* Linux: Set Auto Brightness: manual brightness control is off, rely on ambient;  */
	"VK_BRIGHTNESS_ZERO":  keybd_event.VK_BRIGHTNESS_ZERO,
	"VK_DISPLAY_OFF":      keybd_event.VK_DISPLAY_OFF, /* Linux: display device to off state;  */
	"VK_WWAN":             keybd_event.VK_WWAN,        /* Linux: Wireless WAN (LTE, UMTS, GSM, etc.);  */
	"VK_WIMAX":            keybd_event.VK_WIMAX,
	"VK_RFKILL":           keybd_event.VK_RFKILL,  /* Linux: Key that controls all radios;  */
	"VK_MICMUTE":          keybd_event.VK_MICMUTE, /* Linux: Mute / unmute the microphone;  */
}
