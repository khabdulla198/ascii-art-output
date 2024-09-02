package ascii

import (
	"flag"
	"fmt"
	"strings"
)

func Validation(args []string) (string, string, string, error) {
	file := "standard.txt" //making default as standard
	output := ""

	// TOTO: for loop on args to check for the two wrong cases cases  (--output), (--output=) then print error
	// and exit the program

	fl := flag.NewFlagSet("BannerTool", flag.ExitOnError)
	fl.StringVar(&output, "output", "", "The destination file where the Ascii-Art is saved.")
	fl.Parse(args)

	arg := fl.Args() // Args returns the non-flag arguments.

	for _, arg := range args {
		if arg == "--output" || arg == "--output=" {
			return "", "", "", fmt.Errorf("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		}
	}

	if len(arg) > 2 || len(arg) == 0 {
		//error message for incorrect amount of input
		return "", "", "", fmt.Errorf("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}

	if len(arg) == 2 {
		switch strings.ToLower(arg[1]) {
		// if standard is input, standard.txt will be the file
		case "standard":
			file = "standard.txt"
		// if shadow is input, shadow.txt will become the file
		case "shadow":
			file = "shadow.txt"
		//if thinkertoy is input, thinkertoy will become the file
		case "thinkertoy":
			file = "thinkertoy.txt"
		// error message in case of invalid input
		default:
			return "", "", "", fmt.Errorf("error: only two arguments are allowed [STRING] [OPTIONAL BANNER]\n Banners that are available are: standard, thinkertoy, and shadow")
		}
	}
	wordTofind := arg[0] // declaring word/phrase
	return wordTofind, file, output, nil
}




