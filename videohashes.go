//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"strconv"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"

	"github.com/stashapp/stash/internal/manager"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/oshash"
)

func getFFPaths() (string, string) {
	var paths []string

	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, cwd)
	}

	return ffmpeg.GetPaths(paths)
}

func GeneratePHashValue(encoder ffmpeg.Encoder, videoFile ffmpeg.VideoFile) (*uint64, error) {
	g, err := manager.NewPhashGenerator(videoFile, "")
	if err != nil {
		return nil, err
	}

	// fmt.Printf("[generator] generating phash sprite for %s", g.Info.VideoFile.Path)

	// Generate sprite image offset by 5% on each end to avoid intro/outros
	chunkCount := g.Columns * g.Rows
	offset := 0.05 * g.Info.VideoFile.Duration
	stepSize := (0.9 * g.Info.VideoFile.Duration) / float64(chunkCount)
	var images []image.Image
	for i := 0; i < chunkCount; i++ {
		time := offset + (float64(i) * stepSize)

		options := ffmpeg.SpriteScreenshotOptions{
			Time:  time,
			Width: 160,
		}
		img, err := encoder.SpriteScreenshot(g.Info.VideoFile, options)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	// Combine all of the thumbnails into a sprite image
	if len(images) == 0 {
		return nil, fmt.Errorf("images slice is empty, failed to generate phash sprite for %s", g.Info.VideoFile.Path)
	}
	width := images[0].Bounds().Size().X
	height := images[0].Bounds().Size().Y
	canvasWidth := width * g.Columns
	canvasHeight := height * g.Rows
	montage := imaging.New(canvasWidth, canvasHeight, color.NRGBA{})
	for index := 0; index < len(images); index++ {
		x := width * (index % g.Columns)
		y := height * int(math.Floor(float64(index)/float64(g.Rows)))
		img := images[index]
		montage = imaging.Paste(montage, img, image.Pt(x, y))
	}

	hash, err := goimagehash.PerceptionHash(montage)
	if err != nil {
		return nil, err
	}
	hashValue := hash.GetHash()
	return &hashValue, nil
}

func GeneratePHash(ffmpegPath string, ffprobePath string, videoPath string) (string, int) {
	FFMPEG := ffmpeg.Encoder(ffmpegPath)
	FFProbe := ffmpeg.FFProbe(ffprobePath)

	hexval := ""
	duration := 0

	videoFile, err := FFProbe.NewVideoFile(videoPath, false)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading video file: %s", err.Error()))
		return hexval, duration
	}

	duration = int(videoFile.Duration)

	hash, err := GeneratePHashValue(FFMPEG, *videoFile)
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
