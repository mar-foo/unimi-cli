package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

/* Author: Mario Forzanini https://marioforzanini.com
Date: 04/06/2021
Description: Download and watch video lessons from ariel.unimi.it
Dependencies: dmenu https://tools.suckless.org/dmenu mpv https://mpv.io,
wget, youtube-dl https://www.youtube-dl.org, ffmpeg https://ffmpeg.org */

const inputFile = "webpage.html"

/* Exit codes */
const ExitSucces = 0
const ExitNoDeps = 1
const ExitNoFile = 2
const ExitWrongArg = 3
const ExitNotFound = 4

const usage = `Usage: university-cli [-s SPEED] [OPTIONS]

Available options:
	-h:Display this help message
	-D:Download all videos
	-s:Choose speed of download (possible values: fast, slow),
		has to be set as first option.
`

/*const usage = `Usage: university-cli [-s SPEED] [OPTIONS]

Available options:
	-h:Display this help message
	-c:Check to see which videos aren't downloaded
	-d:Choose a lesson to download
	-D:Download all videos
	-s:Choose speed of download (possible values: fast, slow),
		has to be set as first option.
	-v:Choose a lesson to watch
`*/

type Options struct {
	command  string
	download string
}

type video struct {
	url  string
	name string
}

//main() handles user options:
//university-cli [-s SPEED] [OPTIONS]
//	Available options:
//		-h:Display this help message
//		-c:Check to see which videos aren't downloaded
//		-d:Choose a lesson to download
//		-D:Download all videos
//		-s:Choose speed of download (possible values: fast, slow),
//			has to be set as first option.
//		-v:Choose a lesson to watch
func main() {
	opts := defaultOptions()
	parseOptions(opts, os.Args[1:])
}

/* Error handling */

func noDep(dependency string) {
	if _, err := exec.LookPath(dependency); err != nil {
		log.Fatalln("university-cli: ", dependency, "not found, install it to",
			"proceed")
	}
}

func noFile() {
	log.Fatalln("university-cli: %s not found, use -d to download it.",
		inputFile)
}

func noDownload(filename string) {
	fmt.Println("university-cli: %s already downloaded, nothing to do",
		filename)
	os.Exit(ExitSucces)
}

func noInput(filename string) {
	log.Fatalln("university-cli: No input file found: download the html"+
		"page in a file called", filename, "%s and retry.")
}

//setSpeed(speed string) int sets the speed of the download, accepted
// values are: "fast" and "slow"
func setSpeed(opts *Options, speed string) {
	switch speed {
	case "fast":
		opts.command = "youtube-dl"
	case "slow":
		opts.command = "ffmpeg"
	default:
		log.Fatalln("Speed argument %s not recognised, possible values are: ", speed,
			"'fast' and 'slow'")
	}
}

//checkWebpage(file string) int checks for the existence of the input
//file and find video URLs
func checkWebpage(filename string) []video {
	file, fileError := os.Open(filename)
	if fileError != nil {
		noInput(filename)
	}
	webpage := bufio.NewScanner(file)
	/* videos, names := make([]string, 0), make([]string, 0) */
	videos := make([]video, 0)
	for webpage.Scan() {
		bufLine := webpage.Text()
		urlRegexp := regexp.MustCompile("https?://.*.m3u8")
		nameRegexp := regexp.MustCompile(":[[:alpha:]].*.mp4")
		if hasSubstring(bufLine, "source src=") {
			tempURL := urlRegexp.FindString(bufLine)
			tempName := nameRegexp.FindString(tempURL)[1:]
			tempVideo := &video{
				url:  tempURL,
				name: tempName,
			}
			videos = append(videos, *tempVideo)
		}
	}
	file.Close()
	return videos
}

//download(opts *Options, videos []video) int accept a slice of videos and
//downloads all of them named accordingly
func download(opts *Options, videos []video) {
	if num := len(videos); opts.download == "All" {
		for i := 0; i < num; i++ {
			noDep(opts.command)
			var cmd *exec.Cmd
			if opts.command == "youtube-dl" {
				cmd = exec.Command(opts.command, "--no-check-certificate",
					videos[i].url, "-o", videos[i].name)
			} else {
				// TODO Find a better way to handle multi-option commands
				cmd = exec.Command(opts.command, "-i", videos[i].url, "-n",
					"-preset", "slow", "-c:v", "libx265", "-crf", "31", "-c:a", "aac",
					"-b:a", "64k", "-ac", "1", videos[i].name)
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}
	}
}

//check()
func check() {
	fmt.Print("check")
}

//parseOptions
func parseOptions(opts *Options, allArgs []string) {
	for i := 0; i < len(allArgs); i++ {

		switch arg := allArgs[i]; arg {
		case "-h", "--help":
			fmt.Println(usage)
			os.Exit(0)
		case "-s":
			setSpeed(opts, allArgs[i+1])
			i++
		case "-D":
			opts.download = "All"
			videos := checkWebpage(inputFile)
			download(opts, videos)
		case "-c":
			check()
		default:
			log.Fatalln("Unknown option %s", arg)
		}
	}
}

func defaultOptions() *Options {
	return &Options{
		command:  "youtube-dl",
		download: "All"}
}

func hasSubstring(s string, sub string) bool {
	split := strings.Split(s, sub)
	return len(split) > 1
}
