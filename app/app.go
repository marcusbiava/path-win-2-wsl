package app

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func Run(path string) {
	split := strings.Split(path, ":")
	newPathWithMnt := fmt.Sprintf("/mnt/%s", strings.ToLower(split[0]))
	newPath := strings.ReplaceAll(split[1], string("\\"), "/")
	wslPath := fmt.Sprintf("%s%s", newPathWithMnt, newPath)
	clipboard.WriteAll(wslPath)
	fmt.Println(wslPath)
}
