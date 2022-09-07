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

	ffmpegPath, ffprobePath := internal.GetFFPaths()
	if ffmpegPath == "" || ffprobePath == "" {
		fmt.Println("ffmpeg/ffprobe executables not found")
		return
	}

	result := Result{videoPath: videoPath}

	if err := result.GeneratePHash(ffmpegPath, ffprobePath); err != nil {
		fmt.Println(err)
		return
	}

	if err := result.GenerateOSHash(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n%s\n", result)
}
