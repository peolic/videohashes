package main

import (
	"fmt"
	"strconv"

	"github.com/peolic/videohashes/internal"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/oshash"
	"github.com/stashapp/stash/pkg/hash/videophash"
)

type Result struct {
	videoPath string
	Duration  int
	PHash     string
	OSHash    string
}

func (r *Result) GeneratePHash(ffmpegPath string, ffprobePath string) error {
	FFMPEG := ffmpeg.FFMpeg(ffmpegPath)
	FFProbe := ffmpeg.FFProbe(ffprobePath)

	videoFile, err := FFProbe.NewVideoFile(r.videoPath)
	if err != nil {
		return fmt.Errorf("error reading video file: %s", err.Error())
	}

	r.Duration = int(videoFile.Duration)

	hash, err := videophash.Generate(FFMPEG, videoFile)
	if err != nil {
		return fmt.Errorf("error generating phash: %s", err.Error())
	}

	r.PHash = strconv.FormatUint(*hash, 16)
	return nil
}

func (r *Result) GenerateOSHash() error {
	oshash, err := oshash.FromFilePath(r.videoPath)
	if err != nil {
		return fmt.Errorf("error generating oshash: %s", err)
	}

	r.OSHash = oshash
	return nil
}

func (r Result) String() string {
	buf := ""
	buf += fmt.Sprintf("Duration: %s (%d)\n", internal.FormatDuration(r.Duration), r.Duration)
	buf += fmt.Sprintf("PHash:    %s\n", r.PHash)
	buf += fmt.Sprintf("OSHash:   %s\n", r.OSHash)
	return buf
}
