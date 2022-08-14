package main

import (
	"fmt"
	"os"

	"github.com/peolic/videohashes/internal"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("arg 1: video path")
		return
	}

	_, ffprobePath := internal.GetFFPaths()
	videoPath := args[0]

	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("stat error: %s", err.Error()))
		return
	} else if fileInfo.Mode().IsDir() {
		fmt.Println("file is a directory")
		return
	}

	duration := internal.GetDuration(ffprobePath, videoPath)

	fmt.Println()
	fmt.Printf("Duration: %s (%d)\n", internal.FormatDuration(duration), duration)
	fmt.Println()
}
