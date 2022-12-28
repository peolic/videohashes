package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/peolic/videohashes/internal"
)

func main() {
	videoPath := ""
	ffmpegInstallDir := "./"

	path, err := os.Executable()
	if (err == nil) {
		ffmpegInstallDir = filepath.Dir(path)
	}
	path, err = filepath.Abs(ffmpegInstallDir)
	if (err == nil) {
		ffmpegInstallDir = path
	}

	args := os.Args[1:]
	if len(args) >= 1 {
		videoPath = args[0]
	}

	ffmpegPath, ffprobePath := internal.GetFFPaths(ffmpegInstallDir)
	if ffmpegPath == "" || ffprobePath == "" {
		fmt.Println("acceptable ffmpeg/ffprobe executables not found on path, and could not be installed")
		return
	}

	if videoPath == "" {
		fmt.Println("missing video path")
		return
	}

	if err := internal.ValidFile(videoPath); err != nil {
		fmt.Println(err)
		return
	}

	duration := internal.GetDuration(ffprobePath, videoPath)

	out := fmt.Sprintf("Duration: %s (%d)\n", internal.FormatDuration(duration), duration)
	fmt.Printf("\n%s\n", out)
}
