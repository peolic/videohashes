//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/oshash"
	"github.com/stashapp/stash/pkg/hash/videophash"
)

func getFFPaths() (string, string) {
	var paths []string

	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, cwd)
	}

	return ffmpeg.GetPaths(paths)
}

func GeneratePHash(ffmpegPath string, ffprobePath string, videoPath string) (string, int) {
	FFMPEG := ffmpeg.FFMpeg(ffmpegPath)
	FFProbe := ffmpeg.FFProbe(ffprobePath)

	hexval := ""
	duration := 0

	videoFile, err := FFProbe.NewVideoFile(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading video file: %s", err.Error()))
		return hexval, duration
	}

	duration = int(videoFile.Duration)

	hash, err := videophash.Generate(FFMPEG, videoFile)
	if err != nil {
		fmt.Println(fmt.Errorf("error generating phash: %s", err.Error()))
		return hexval, duration
	}

	hexval = strconv.FormatUint(*hash, 16)
	return hexval, duration
}

func GenerateOSHash(videoPath string) string {
	oshash, err := oshash.FromFilePath(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error generating oshash: %s", err))
		return ""
	}

	return oshash
}

func formatDuration(duration int) string {
	hours := duration / 3600
	minutes := (duration % 3600) / 60
	seconds := duration % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("arg 1: video path")
		return
	}

	ffmpegPath, ffprobePath := getFFPaths()
	videoPath := args[0]

	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("stat error: %s", err.Error()))
		return
	} else if fileInfo.Mode().IsDir() {
		fmt.Println("file is a directory")
		return
	}

	phash, duration := GeneratePHash(ffmpegPath, ffprobePath, videoPath)
	oshash := GenerateOSHash(videoPath)

	fmt.Println()
	fmt.Printf("Duration: %s (%d)\n", formatDuration(duration), duration)
	fmt.Printf("PHash:    %s\n", phash)
	fmt.Printf("OSHash:   %s\n", oshash)
	fmt.Println()
}
