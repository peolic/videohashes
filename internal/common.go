package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/file"
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

func GetFFPaths(targetDir string) (string, string) {
	var paths = []string{targetDir}

	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, cwd)
	}
    return GetFFMPEG(targetDir, paths)
}

func GetFFMPEG(targetDir string, paths []string) (string, string) {
	if (targetDir != "") {
		var ctx = context.Background()
		fullpaths := append([]string{targetDir}, paths...)

		ffmpegPath, ffprobePath := ffmpeg.GetPaths(fullpaths)

		if ffmpegPath == "" || ffprobePath == "" {
			if err := ffmpeg.Download(ctx, targetDir); err != nil {
				msg := `Unable to locate / automatically download FFMPEG
Check the readme for download links.
The FFMPEG and FFProbe binaries should be placed in %s
The error was: %s
`
				fmt.Printf(msg, targetDir, err)
				return "", ""
			}
		}
		return ffmpeg.GetPaths(fullpaths)
	}
	return "",""
}


func GetDuration(ffprobePath string, videoPath string) int {
	FFProbe := ffmpeg.FFProbe(ffprobePath)

	videoProbe, err := FFProbe.NewVideoFile(videoPath)
	if err != nil {
		fmt.Println(fmt.Errorf("error reading video file: %s", err.Error()))
		return 0
	}

	return int(videoProbe.FileDuration)
}

func FormatDuration(duration int) string {
	hours := duration / 3600
	minutes := (duration % 3600) / 60
	seconds := duration % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// Based on https://github.com/stashapp/stash/blob/8b59a3b01418/pkg/file/video/scan.go#L17
func ProbeResultToVideoFile(videoProbe *ffmpeg.VideoFile, videoPath string) (*file.VideoFile, error) {
	container, err := ffmpeg.MatchContainer(videoProbe.Container, videoPath)
	if err != nil {
		return nil, fmt.Errorf("matching container for %q: %w", videoPath, err)
	}

	videoFile := &file.VideoFile{
		BaseFile: &file.BaseFile{
			Path: videoPath,
		},
		Format:      string(container),
		VideoCodec:  videoProbe.VideoCodec,
		AudioCodec:  videoProbe.AudioCodec,
		Width:       videoProbe.Width,
		Height:      videoProbe.Height,
		Duration:    videoProbe.FileDuration,
		FrameRate:   videoProbe.FrameRate,
		BitRate:     videoProbe.Bitrate,
		Interactive: false,
	}

	return videoFile, nil
}
