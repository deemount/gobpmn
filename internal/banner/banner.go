package banner

import (
	"fmt"
	"strings"
)

// ANSI-colors for terminal output in rainbow
// The colors are defined as ANSI escape codes.
var colors = []string{
	"\033[38;5;196m", // red
	"\033[38;5;202m", // orange
	"\033[38;5;226m", // yellow
	"\033[38;5;46m",  // green
	"\033[38;5;51m",  // cyan
	"\033[38;5;21m",  // blue
	"\033[38;5;201m", // magenta
}

const reset = "\033[0m"

const b = `
                  888
                  888
                  888
 .d88b.   .d88b.  88888b.  88888b.  88888b.d88b.  88888b.
d88P"88b d88""88b 888 "88b 888 "88b 888 "888 "88b 888 "88b
888  888 888  888 888  888 888  888 888  888  888 888  888
Y88b 888 Y88..88P 888 d88P 888 d88P 888  888  888 888  888
 "Y88888  "Y88P"  88888P"  88888P"  888  888  888 888  888
     888                   888
Y8b d88P                   888
 "Y88P"                    888
`

func PrintRainbowBanner() {
	lines := strings.Split(b, "\n")
	for i, line := range lines {
		color := colors[i%len(colors)]
		fmt.Println(color + line + reset)
	}
}
