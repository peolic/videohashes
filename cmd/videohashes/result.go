package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/peolic/videohashes/internal"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/hash/md5"
	"github.com/stashapp/stash/pkg/hash/oshash"
	"github.com/stashapp/stash/pkg/hash/videophash"
)

type Result struct {
	videoPath string `json:"-"`
	Duration  int    `json:"duration"`
	PHash     string `json:"phash"`
	OSHash    string `json:"oshash"`
	MD5       string `json:"md5,omitempty"`
}

func (r *Result) GeneratePHash(ffmpegPath string, ffprobePath string) error {
	FFMPEG := ffmpeg.NewEncoder(ffmpegPath)
	FFProbe := ffmpeg.NewFFProbe(ffprobePath)

	videoProbe, err := FFProbe.NewVideoFile(r.videoPath)
	if err != nil {
		return fmt.Errorf("error reading video file: %s", err.Error())
	}

	videoFile, err := internal.ProbeResultToVideoFile(videoProbe, r.videoPath)
	if err != nil {
		return fmt.Errorf("error coverting probe result: %s", err.Error())
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

func (r *Result) GenerateMD5() error {
	md5, err := md5.FromFilePath(r.videoPath)
	if err != nil {
		return fmt.Errorf("error generating md5: %s", err)
	}

	r.MD5 = md5
	return nil
}

func (r Result) String() string {
	buf := ""
	buf += fmt.Sprintf("Duration: %s (%d)\n", internal.FormatDuration(r.Duration), r.Duration)
	buf += fmt.Sprintf("PHash:    %s\n", r.PHash)
	buf += fmt.Sprintf("OSHash:   %s\n", r.OSHash)
	if r.MD5 != "" {
		buf += fmt.Sprintf("MD5:      %s\n", r.MD5)
	}
	return buf
}

func (r Result) JSON() (string, error) {
	out, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
