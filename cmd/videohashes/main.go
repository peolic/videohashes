package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/peolic/videohashes/internal"
)

func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] video\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	videoPath := ""
	calcMD5 := false
	jsonOut := false
	ffmpegInstallDir := "./"

	path, err := os.Executable()
	if (err == nil) {
		ffmpegInstallDir = filepath.Dir(path)
	}
	path, err = filepath.Abs(ffmpegInstallDir)
	if (err == nil) {
		ffmpegInstallDir = path
	}

	flag.StringVar(&videoPath, "video", "", "path to video file")
	flag.BoolVar(&calcMD5, "md5", false, "calculate md5 checksum as well")
	flag.BoolVar(&jsonOut, "json", false, "output in json format")
	flag.StringVar(&ffmpegInstallDir, "ffmpeg", ffmpegInstallDir, "where to install ffmpeg if needed")
	flag.Usage = myUsage
	flag.Parse()

	ffmpegPath, ffprobePath := internal.GetFFPaths(ffmpegInstallDir)
	if ffmpegPath == "" || ffprobePath == "" {
		fmt.Println("acceptable ffmpeg/ffprobe executables not found on path, and could not be installed")
		return
	}

	if videoPath == "" {
		videoPath = flag.Arg(0)
	}

	if videoPath == "" {
		fmt.Println("missing video path")
		flag.Usage()
		return
	}

	if err := internal.ValidFile(videoPath); err != nil {
		fmt.Println(err)
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

	if calcMD5 {
		if err := result.GenerateMD5(); err != nil {
			fmt.Println(err)
			return
		}
	}

	if jsonOut {
		out, err := result.JSON()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(out)
		return
	}

	fmt.Printf("\n%s\n", result)
}
