package keys

import (
	"fmt"
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

type Key struct {
	Name     string `json:"key"`
	kbObject keybd_event.KeyBonding
}

func NewKey() (*Key, error) {
	key := &Key{}
	kb, err := keybd_event.NewKeyBonding()
	if err == nil {
		key.kbObject = kb
		if runtime.GOOS == "linux" {
			time.Sleep(2 * time.Second)
		}
	}
	return key, err
}

func (key *Key) Press() error {
	if pressedKey, ok := keyboardMap[key.Name]; ok {
		key.kbObject.SetKeys(pressedKey)
		err := key.kbObject.Launching()
		return err
	} else {
		return fmt.Errorf("unsupported key \"%v\"", key.Name)
	}
}
