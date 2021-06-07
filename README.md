# university-cli
Command Line Interface to ariel.unimi.it

# Installation
The program is just a simple `bash` script (will try to make it POSIX compliant in the future) so it only works on GNU/Linux and maybe macOS (not tested).
I rewrote the script in [go](https://golang.org) to make it multi-platform and to practise with the language. The `bash` script has been moved to **university-cli.sh** while compiled binaries have been moved into **bin/**.

Install the [dependencies](#dependencies) then clone the repo

~~~ sh
$ git clone https://github.com/mar-foo/university-cli
$ cd university-cli
~~~

## Linux / MacOs
Run

~~~ sh
$ ./install.sh
~~~

by default the program is copied to ~/.local/bin, to add it to your $PATH run

~~~ sh
$ export PATH="$HOME/.local/bin:${PATH}"
~~~

Your shell looks for programs in the list of directories stored in this variable, as soon as you close the terminal you are working in the value of the $PATH variable will be overwritten by the default value; to overcome this issue you have to add this configuration to `$HOME/.bashrc` if you are running `bash` or `$HOME/.zshrc` if you are running `zsh` as your default shell. To do so copy the following command:

~~~ sh
$ echo 'export PATH="$HOME/.local/bin:${PATH}"' >> ~/.bashrc # If running bash
$ echo 'export PATH="$HOME/.local/bin:${PATH}"' >> ~/.zshrc # If running zsh
~~~

## Windows (NOT tested)
Having a go version of the program makes it possible to cross compile it for Windows, I don't have any Windows machine at the moment so I can't test wether it works or not.

### If you have `git` downloaded
Install the [dependencies](#dependencies) then clone the repo

~~~ sh
> git clone https://github.com/mar-foo/university-cli
> cd university-cli
~~~

Move the executable in a directory in your [PATH](https://stackoverflow.com/questions/45072617/default-values-of-path-variables-in-windows-10), you will find it in `bin/windows/university-cli.exe`.
<!-- Alternatively if you have `GNU Make` installed you can simply run:

~~~ sh
> make install
~~~ -->

### If you **don't** have `git` downloaded
Copy the executable directly from this webpage: you can find it in `bin/windows/university-cli.exe` and save it in a directory in your [PATH](https://stackoverflow.com/questions/45072617/default-values-of-path-variables-in-windows-10).

Consider downloading `git` because you can use it to automatically [update](#update) to the newest version.

## From source (All platforms)
If you want to help with the development yourself you will need to be able to compile the source code into an executable, to do so i use [go](https://golang.org): `cd` into the installation directory and run:

~~~ sh
$ go build university-cli.go
~~~

This will create the new executable in the directory root.

## Dependencies
- [dmenu](https://tools.suckless.org/dmenu) to choose videos to download/stream
- [mpv](https://mpv.io) to stream videos
- [youtube-dl](https://youtube-dl.org) to download videos
- [ffmpeg](https://ffmpeg.org) to encode videos

# Usage
At the moment you **need** to have the html source code for the webpage of the course downloaded to a file called *webpage.html*, to download it look up the documentation for your web browser.

**university-cli** will look for the *webpage* file in the current directory, so first you have to create a directory in which you want to store the videos and then cd into it, e.g. to download videos in ~/Documents/University/Videos do:

~~~ sh
$ mkdir -p ~/Documents/University/Videos ; cd ~/Documents/University/Videos # make the directory if it is not there
$ # make sure you have the webpage file in this directory
$ university-cli -D
~~~

Full list of command line options:
- `university-cli -c` check for videos that were not downloaded
- `university-cli -d` prompt for a video to download (uses `dmenu`) (`university-cli.sh` only)
- `university-cli -D` download all of the .mp4 files found on the webpage
- `university-cli -s [SPEED]` set speed of the download (has to be set as first option to work), possible values are 'fast' or 'slow'
- `university-cli -v` prompt for a video to stream (uses `mpv`) (`university-cli.sh` only)
- `university-cli -h` prints usage information

Regarding the `-s` option: the 'fast' method uses youtube-dl's default encoding, output files will be bigger in size but will be encoded faster; on the other hand the 'slow' method uses `ffmpeg` with the options specified [here](https://wiki.studentiunimi.it/guida:scaricare_videolezioni_ariel) which will result in smaller files but will require more time to encode. If no option is specified the default value is 'fast'.

# Example usage
Using the same example directory ~/Documents/University/Videos, to view a video lesson you would run:

~~~ sh
$ cd ~/Documents/University/Videos
$ university-cli -v # Prompt for a lesson to stream
~~~
by default it will look for a local file that matches the chosen name in the current directory, if it is not there it will stream it.

to download it instead you would do:

~~~ sh
$ university-cli -s 'fast' -d
~~~
this promts for a lesson to download using the 'fast' method.

# Update
To keep the installation up to date periodically `cd` into the directory where you cloned the repository and run:

~~~ sh
$ git pull
~~~

And repeat the steps you followed when [installing](#Installation).

# Status
The project is still immature, in the future this wants to be an ncurses/cli interface to browse Ariel.
At the moment the bash script `university-cli.sh` supports more options, but I'm definitively switching to the GoLang version so I'll implement all of those features in the future.

## TODO
- Test for different courses: teachers might store their lessons on different sites which might not be supported by youtube-dl
- Fix issue with spaces: if the filename contains a space then the program might act weird
- Add a test to check if every lesson has been installed by `-D` option

# Contributing
I wrote this little program for myself, if you need/want new features patches are welcome, if you find any bugs file an issue on github and I'll try to figure it out.

# Copyright
This project is not endorsed or approved by the ["Universita degli Studi di Milano"](https://unimi.it), all of the content stored on [ariel](https://ariel.unimi.it) is protected by [copyright](https://ariel.unimi.it/documenti/copyright) as stated by unimi and should be used according to the rules.
