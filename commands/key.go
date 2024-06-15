package commands

import (
	"os"
	"path"

	"github.com/Mahmoud-Khaled-FS/clickterm/utils"
)

type KeyCommand struct {
	Key         string
	KeyPath     string
	ShouldClear bool
}

func (k *KeyCommand) Set() {
	key := utils.EncodeStr(k.Key)
	os.MkdirAll(path.Dir(k.KeyPath), 0755)
	err := os.WriteFile(k.KeyPath, []byte(key), 0755)
	if err != nil {
		panic(err)
	}
}

func (k *KeyCommand) Clear() {
	os.Remove(k.KeyPath)
}

func (k *KeyCommand) Check() bool {
	_, err := os.Stat(k.KeyPath)
	if err != nil {
		return false
	}
	if _, err = utils.DecodeStr(k.Key); err != nil {
		return false
	}
	return true
}
