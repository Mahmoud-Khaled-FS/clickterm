package cmd

import (
	"fmt"

	"github.com/Mahmoud-Khaled-FS/clickterm/commands"
	"github.com/spf13/cobra"
)

func keyCmd() *cobra.Command {
	key := commands.KeyCommand{
		KeyPath: "./data/token.key",
	}
	var setKeyCmd = &cobra.Command{
		Use:   "key",
		Short: "Set or check the api token",
		Long:  "Manage your ClickUp tasks efficiently from the command line with our versatile tool.",
		Run: func(cmd *cobra.Command, args []string) {
			if key.Key != "" {
				key.Set()
				return
			}
			if key.ShouldClear {
				key.Clear()
				return
			}
			if key.Check() {
				fmt.Println("Key is set!")
				return
			}
			fmt.Println("key is not set!")
		},
	}
	setKeyCmd.Flags().StringVarP(&key.Key, "set", "s", "", "Set Api token")
	setKeyCmd.Flags().BoolVarP(&key.ShouldClear, "clear", "c", false, "Remove api key")
	return setKeyCmd
}
