package utils

import (
	"fmt"
	"os"
	"unicode/utf8"

	"golang.org/x/term"
)

func SupportsColour() bool {
    return term.IsTerminal(int(os.Stdout.Fd()))
}

func SupportsUnicode(unicode string) bool {
	return utf8.ValidString(fmt.Sprint(unicode))
}