# university-cli
Command Line Interface to ariel.unimi.it

# Install
Install the [dependencies](#Dependencies), clone the repo

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
At the moment you need to have the html source code for the webpage of the course downloaded to a file called *webpage.html*, to download it look up the documentation for your browser.

# Status
The project is still immature, in the future this wants to be an ncurses/cli interface to browse Ariel.
At the moment watching and downloading videos is the only thing supported.

# Copyright
This project is not endorsed or approved by the ["Universita degli Studi di Milano"](https://unimi.it), all of the content stored on [ariel](https://ariel.unimi.it) is protected by copyright as stated by [unimi](https://ariel.unimi.it/documenti/copyright) and should be used according to the rules.
