package internal

import (
	"fmt"
	"os"

	"github.com/stashapp/stash/pkg/ffmpeg"
)

func ValidFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("stat error: %s", err.Error())
	} else if fileInfo.Mode().IsDir() {
		return fmt.Errorf("file is a directory")
	}
	return nil
}

func GetFFPaths() (string, string) {
	var paths []string

	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, cwd)
	}

	return ffmpeg.GetPaths(paths)
}

func GetDuration(ffprobePath string, videoPath string) int {
	FFProbe := ffmpeg.FFProbe(ffprobePath)

	videoFile, err := FFProbe.NewVideoFile(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading video file: %s", err.Error()))
		return 0
	}

	return int(videoFile.Duration)
}

func FormatDuration(duration int) string {
	hours := duration / 3600
	minutes := (duration % 3600) / 60
	seconds := duration % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
