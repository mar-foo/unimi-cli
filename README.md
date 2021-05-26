# university-cli
Command Line Interface to ariel.unimi.it

# Installation
The program is just a simple bash script (will try to make it POSIX compliant in the future) so it only works on GNU/Linux and maybe macOS (not tested).
Install the [dependencies](#dependencies) then clone the repo

~~~ sh
$ git clone https://github.com/mar-foo/university-cli
$ cd university-cli
$ make install
~~~

by default the script is copied to ~/.local/bin, to add it to your $PATH run

~~~ sh
$ export PATH="$HOME/.local/bin:${PATH}"
~~~

To have your $PATH automatically updated add the command to your shell's configuration file (~/.bashrc, ~/.zshrc ...); if you don't know which shell you're running run

~~~ sh
$ echo $0
~~~

## Dependencies
- [dmenu](https://tools.suckless.org/dmenu) to choose videos to download/stream
- [mpv](https://mpv.io) to stream videos
- [youtube-dl](https://youtube-dl.org) to download videos

# Usage
At the moment you need to have the html source code for the webpage of the course downloaded to a file called *webpage.html*, to download it look up the documentation for your web browser.

**university-cli** will look for the *webpage* file in the current directory, so first you have to create a directory in which you want to store the videos and then cd into it, e.g. to download videos in ~/Documents/University/Videos do:

~~~ sh
$ mkdir -p ~/Documents/University/Videos ; cd ~/Documents/University/Videos # make the directory if it is not there
$ # make sure you have the webpage file in this directory
$ university-cli -D
~~~

Full list of command line options:
- `university-cli -c` check for videos that were not downloaded
- `university-cli -d` prompt for a video to download (uses `dmenu`)
- `university-cli -D` download all of the .mp4 files found on the webpage
- `university-cli -v` prompt for a video to stream (uses `mpv`)
- `university-cli -h` prints usage information

# Example usage
Using the same example directory ~/Documents/University/Videos, to view a video lesson you would run:

~~~ sh
$ cd ~/Documents/University/Videos
$ university-cli -v # Prompt for a lesson to stream
~~~
by default it will look for a local file that matches the chosen name in the current directory, if it is not there it will stream it.

# Status
The project is still immature, in the future this wants to be an ncurses/cli interface to browse Ariel.
At the moment watching and downloading videos is the only thing supported.

## TODO
- Test for different courses: teachers might store their lessons on different sites which might not be supported by youtube-dl
- Fix issue with spaces: if the filename contains a space then the program might act weird
- Add a test to check if every lesson has been installed by `-D` option

# Contributing
I wrote this little program for myself, if you need/want new features patches are welcome, if you find any bugs file an issue on github and I'll try to figure it out.

# Copyright
This project is not endorsed or approved by the ["Universita degli Studi di Milano"](https://unimi.it), all of the content stored on [ariel](https://ariel.unimi.it) is protected by [copyright](https://ariel.unimi.it/documenti/copyright) as stated by unimi and should be used according to the rules.
