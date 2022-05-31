package printer

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Colors functions for printing.
var (
	Yellow = color.New(color.FgYellow).SprintfFunc()
	Red    = color.New(color.FgRed).SprintfFunc()
	Green  = color.New(color.FgGreen).SprintfFunc()
)

func Print(tag, text string) {
	fmt.Fprintf(color.Output, "%s: %s \n", Yellow(tag), text)
}

func Fatal(tag, text string, err error) {
	fmt.Fprintf(color.Output, "%s: %s \n", Yellow(tag), Red(text+": "+err.Error()))
	os.Exit(1)
}
