package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	sourceFlag = flag.String("s", "", "source videofile name")
	outFlag    = flag.String("o", "", "output videofile name")
)

type VideoFile struct {
	src string
	out string
}

func main() {
	args := parseArguments()
	ffmpegCmd := exec.Command("ffmpeg", "-i", args.src, "-vcodec", "libx265", "-crf", "28", args.out)
	ffmpegOut, err := ffmpegCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %v", ffmpegOut)
}

func parseArguments() (file VideoFile) {
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()
	file.src = *sourceFlag
	file.out = *outFlag
	return
}
