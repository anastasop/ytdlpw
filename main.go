package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
)

var (
	quiet           = flag.Bool("q", false, "silent mode")
	writeInfo       = flag.Bool("i", false, "write info file")
	onlyAudio       = flag.Bool("a", false, "download only audio")
	maxWidth        = flag.Int("w", 1280, "choose video with at least this video width")
	fixTitle        = flag.Bool("e", false, "set the video title based on the URL path. For cases like .../name/index.mpd")
	customFormat    = flag.String("f", "", "custom output name format for yt-dlp")
	continueOnError = flag.Bool("c", false, "continue fetching urls even if some of them failed")
)

func usage() {
	fmt.Fprintf(os.Stderr, `usage: ytdlpw -q -i -a -w <width> -e -f <format> -c urls...

ytdlpw is a wrapper for yt-dlp with reasonable defaults for filenames. It can be used to create a media library.
Flags:
`)
	flag.PrintDefaults()
	os.Exit(2)
}

func fetch(mediaURL string) error {
	title := "%(title)s.%(ext)s"
	if *fixTitle {
		u, err := url.Parse(mediaURL)
		if err != nil {
			log.Fatalf("failed to parse: %s", mediaURL)
		}
		title = path.Base(path.Dir(path.Clean(u.Path))) + ".%(ext)s"
	}

	format := fmt.Sprintf("b[width<=%d]/bv*[width<=%d]+ba", *maxWidth, *maxWidth)
	if *onlyAudio {
		format = "bestaudio"
	} else if *customFormat != "" {
		format = *customFormat
	}

	var options []string
	if *quiet {
		options = append(options, "--quiet")
		options = append(options, "--no-warnings")
	}
	if *writeInfo {
		options = append(options, "--write-info-json")
	}
	options = append(options, "--abort-on-error")
	options = append(options, "--restrict-filenames")
	options = append(options, "--windows-filenames")
	options = append(options, "--no-split-chapters")
	options = append(options, "--no-overwrites")
	options = append(options, "-S")
	options = append(options, "ext")
	options = append(options, "--format")
	options = append(options, format)
	options = append(options, "--output")
	options = append(options, title)
	options = append(options, mediaURL)

	cmd := exec.Command("yt-dlp", options...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	for _, u := range flag.Args() {
		if err := fetch(u); err != nil {
			log.Printf("error %s: %v\n", u, err)
		}
		if !*continueOnError {
			os.Exit(1)
		}
	}
}
