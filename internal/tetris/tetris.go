package tetris

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func resetTerminal() {
	fmt.Print("\033[0m")  // Reset colors
	fmt.Print("\033[?25h") // Show cursor
}

func Start() {
	//Set raw mode, disable on refer
	oldTerm, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldTerm) 

	defer resetTerminal() // Ensure reset when the program exits

	for {
		var b = make([]byte, 1) 
		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Failed to read input. Ignoring: ", err)
			continue
		}

		key := b[0]
		if key == 'q' {
			fmt.Println("Quitting...")
			break
		}
	}

	// defer term.Restore(int(os.Stdin.Fd()), oldState) // Restore the terminal state when done
}