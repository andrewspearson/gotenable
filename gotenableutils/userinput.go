package gotenableutils

import (
	"fmt"
	"github.com/andrewspearson/gotenable"
	"golang.org/x/term"
	"os"
)

// UserInput captures user input and has configurable echo
func UserInput(prompt string, echo bool) string {
	var data string
	fmt.Print(prompt)
	if echo == true {
		_, err := fmt.Scanln(&data)
		ErrFatal(err)
	} else if echo == false {
		dataBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		ErrFatal(err)
		fmt.Println()
		data = string(dataBytes)
	} else {
		gotenable.Log.Logger.Fatalln(`[FATAL] Failed to determine if user input should echo to terminal.`)
	}
	return data
}
