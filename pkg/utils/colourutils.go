package utils

import (
	"os"

	"golang.org/x/term"
)

func SupportsColour() bool {
    return term.IsTerminal(int(os.Stdout.Fd()))
}
