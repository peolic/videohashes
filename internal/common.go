package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
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
	var ffmpegPath, ffprobePath string

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ffmpegPath = ffmpeg.ResolveFFMpeg(exPath, cwd)
	ffprobePath = ffmpeg.ResolveFFProbe(exPath, cwd)

	return ffmpegPath, ffprobePath
}

func GetDuration(ffprobePath string, videoPath string) int {
	FFProbe := ffmpeg.NewFFProbe(ffprobePath)

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
func ProbeResultToVideoFile(videoProbe *ffmpeg.VideoFile, videoPath string) (*models.VideoFile, error) {
	container, err := ffmpeg.MatchContainer(videoProbe.Container, videoPath)
	if err != nil {
		return nil, fmt.Errorf("matching container for %q: %w", videoPath, err)
	}

	videoFile := &models.VideoFile{
		BaseFile: &models.BaseFile{
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
