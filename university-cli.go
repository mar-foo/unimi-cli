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
  -h,  --help:         Display this help message
  -c,  --check:        Check to see which videos aren't downloaded
  -D,  --download-all: Download all videos
  -s,  --speed:        Choose speed of download (possible values: fast, slow),
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

//check() is run when the -c option is given, it checks if every videos has been
//downloaded based on their name
func check(videos []video) {
	var exitCode int
	for i := 0; i < len(videos); i++ {
		if !fileExists(videos[i].name) {
			fmt.Println("university-cli: file", videos[i].name, "not found.")
			exitCode = ExitNoFile
		}
	}
	if exitCode == 0 {
		fmt.Println("Found every video, exiting.")
	}
	os.Exit(exitCode)
}

//checkWebpage(file string) int checks for the existence of the input
//file named "filename" and returns an array of videos it found
func checkWebpage(filename string) []video {
	file, fileError := os.Open(filename)
	if fileError != nil {
		noFile(filename)
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
			tempName := nameRegexp.FindString(tempURL)
			if tempName != "" {
				/* If the filename contains "/" youtube-dl and ffmpeg will create
				undesired subdirectories */
				tempName := strings.Split(tempName[1:], "/")
				tempVideo := &video{
					url:  tempURL,
					name: tempName[len(tempName)-1],
				}
				videos = append(videos, *tempVideo)
			}
		}
	}
	file.Close()
	return videos
}

//defaultOptions() returns a pointer to an array of Options set to the default
//values
func defaultOptions() *Options {
	return &Options{
		command:  "youtube-dl", //fast download method
		download: "All"}        //Download all the videos by default
}

//download(opts *Options, videos []video) int accept a slice of videos and
//downloads them named accordingly
func download(opts *Options, videos []video) {
	if num := len(videos); opts.download == "All" {
		for i := 0; i < num; i++ {
			noDep(opts.command)
			var cmd *exec.Cmd
			if opts.command == "youtube-dl" {
				cmd = exec.Command(opts.command, "--no-check-certificate",
					"--no-overwrite", videos[i].url, "-o", videos[i].name)
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

//fileExists(filename string) returns true if the file called "filename" exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//hasSubstring(s string, sub string) returns true if s contains at least one
//occurrence of sub
func hasSubstring(s string, sub string) bool {
	split := strings.Split(s, sub)
	return len(split) > 1
}

//noDep(dependency string) ) exits if dependency is not found in PATH
func noDep(dependency string) {
	if _, err := exec.LookPath(dependency); err != nil {
		log.Fatalln("university-cli: ", dependency, "not found, install it to",
			"proceed")
	}
}

//noDownload(filename string) checks if filename has already been downloaded
func noDownload(filename string) {
	fmt.Println("university-cli:", filename, "already downloaded, nothing to do")
	os.Exit(ExitSucces)
}

//noFile(filename string) exits if filename is not found in the current
//directory
func noFile(filename string) {
	log.Fatalln("university-cli:", filename, "not found.")
}

//parseOptions(opts *Options, allArgs []string) parses options
//in AllArgs and fills the array of Options pointed by opts accordingly
func parseOptions(opts *Options, allArgs []string) {
	for i := 0; i < len(allArgs); i++ {
		switch arg := allArgs[i]; arg {
		case "-h", "--help":
			fmt.Println(usage)
			os.Exit(0)
		case "-s", "--speed":
			setSpeed(opts, allArgs[i+1])
			i++
		case "-D", "--download-all":
			opts.download = "All"
			videos := checkWebpage(inputFile)
			download(opts, videos)
		case "-c", "--check":
			videos := checkWebpage(inputFile)
			check(videos)
		default:
			fmt.Println("Unknown option", arg)
			fmt.Println(usage)
			os.Exit(ExitWrongArg)
		}
	}
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
		log.Fatalln("Speed argument", speed, "not recognised, possible values are:",
			"'fast' and 'slow'")
	}
}
