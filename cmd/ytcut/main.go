package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

func parseArgs() (url, start, end string) {
	flag.StringVar(&url, "u", "", "youtube video url")
	flag.StringVar(&start, "s", "*00:00:00", "start time 99:59:59.99")
	flag.StringVar(&end, "e", "00:00:01", "ending time 99:59:59.99")
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	flag.Parse()
	return
}

func makeTimeInterval(start, end string) string {
	return "*" + strings.Join([]string{start, end}, "-")
}

func main() {
	url, start, end := parseArgs()
	cmd := exec.Command(
		"yt-dlp",
		"--download-sections",
		makeTimeInterval(start, end),
		"--force-keyframes-at-cuts",
		"--ppa",
		"ffmpeg:-http_persistent 0",
		url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
