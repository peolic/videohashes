package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/peolic/videohashes/internal"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/oshash"
	"github.com/stashapp/stash/pkg/hash/videophash"
)

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

	phash, duration := GeneratePHash(ffmpegPath, ffprobePath, videoPath)
	if phash == "" {
		return
	}
	oshash := GenerateOSHash(videoPath)
	if oshash == "" {
		return
	}

	fmt.Println()
	fmt.Printf("Duration: %s (%d)\n", internal.FormatDuration(duration), duration)
	fmt.Printf("PHash:    %s\n", phash)
	fmt.Printf("OSHash:   %s\n", oshash)
	fmt.Println()
}
