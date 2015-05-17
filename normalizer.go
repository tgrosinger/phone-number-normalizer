package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getFromClipboard() string {
	var out bytes.Buffer
	xsel := exec.Command("xsel")
	xsel.Stdout = &out
	err := xsel.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return out.String()
}

func writeToClipboard(s string) {
	xsel := exec.Command("xclip", "-selection", "clipboard")
	xsel.Stdin = strings.NewReader(s)
	err := xsel.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	replacer := strings.NewReplacer(
		"(", "",
		")", "",
		"-", "",
		".", "",
		" ", "",
	)

	previous := getFromClipboard()

	for {
		result := getFromClipboard()
		if result == previous || result == "" {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		previous = result

		if strings.HasPrefix(result, "1") {
			result = result[1:]
		}
		if strings.HasPrefix(result, "+1") {
			result = result[2:]
		}

		result = replacer.Replace(result)

		if len(result) != 10 {
			fmt.Println("Not a phone number:", result)
			continue
		}

		result = fmt.Sprintf("+1%s", result)
		writeToClipboard(result)
		fmt.Println("Converted", previous, "into", result)
	}
}
