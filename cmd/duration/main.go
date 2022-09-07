package main

import (
	"fmt"
	"os"

	"github.com/peolic/videohashes/internal"
)

func main() {
	videoPath := ""

	args := os.Args[1:]
	if len(args) >= 1 {
		videoPath = args[0]
	}

	if videoPath == "" {
		fmt.Println("missing video path")
		return
	}

	if err := internal.ValidFile(videoPath); err != nil {
		fmt.Println(err)
		return
	}

	_, ffprobePath := internal.GetFFPaths()
	if ffprobePath == "" {
		fmt.Println("ffprobe executable not found")
		return
	}

	duration := internal.GetDuration(ffprobePath, videoPath)

	out := fmt.Sprintf("Duration: %s (%d)\n", internal.FormatDuration(duration), duration)
	fmt.Printf("\n%s\n", out)
}
