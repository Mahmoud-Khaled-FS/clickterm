package commands

import (
	"fmt"
	"os"
	"path"

	"github.com/Mahmoud-Khaled-FS/clickterm/clickup"
	"github.com/Mahmoud-Khaled-FS/clickterm/utils"
)

type KeyCommand struct {
	Key          string
	KeyPath      string
	ShouldClear  bool
	ShouldVerify bool
}

func (k *KeyCommand) Set() {
	if k.ShouldVerify {
		k.verifyUser()
	}
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
 	encodedKey, err := os.ReadFile(k.KeyPath)
	if err != nil {
		panic(err)
	}
	k.Key, err = utils.DecodeStr(string(encodedKey))
	if err != nil {
		return false
	}
	if k.ShouldVerify {
		k.verifyUser()
	}
	return true
}

func (k *KeyCommand) verifyUser() {
	user, err := clickup.New(k.Key).GetUser()
	if err != nil {
		fmt.Println("Invalid token!")
		os.Exit(1)
	}
	fmt.Printf("Hello %s\n", user.Username)
}
